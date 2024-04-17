package statetracker

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/lavanet/lava/app"
	"github.com/lavanet/lava/protocol/chainlib"
	"github.com/lavanet/lava/protocol/chaintracker"
	updaters "github.com/lavanet/lava/protocol/statetracker/updaters"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/utils/rand"
	"github.com/lavanet/lava/utils/sigs"
	"github.com/spf13/cobra"
)

const (
	FlagTimeout                 = "timeout"
	FlagValue                   = "value"
	FlagEventName               = "event"
	FlagBreak                   = "break"
	FlagHasAttributeName        = "has-attribute"
	FlagShowAttributeName       = "show-attribute"
	FlagDisableInteractiveShell = "disable-interactive"
)

func eventsLookup(ctx context.Context, clientCtx client.Context, blocks, fromBlock int64, eventName, value string, shouldBreak bool, hasAttributeName string, showAttributeName string, disableInteractive bool) error {
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	resultStatus, err := clientCtx.Client.Status(ctx)
	if err != nil {
		return err
	}
	latestHeight := resultStatus.SyncInfo.LatestBlockHeight
	if latestHeight < blocks {
		return utils.LavaFormatError("requested blocks is bigger than latest block height", nil, utils.Attribute{Key: "requested", Value: blocks}, utils.Attribute{Key: "latestHeight", Value: latestHeight})
	}
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	readEventsFromBlock := func(blockFrom int64, blockTo int64, hash string) {
		for block := blockFrom; block < blockTo; block++ {
			brp, err := updaters.TryIntoTendermintRPC(clientCtx.Client)
			if err != nil {
				utils.LavaFormatFatal("invalid blockResults provider", err)
			}
			blockResults, err := brp.BlockResults(ctx, &block)
			if err != nil {
				utils.LavaFormatError("invalid blockResults status", err)
				return
			}
			for _, event := range blockResults.BeginBlockEvents {
				checkEventForShow(eventName, event, hasAttributeName, value, block, showAttributeName)
			}
			transactionResults := blockResults.TxsResults
			for _, tx := range transactionResults {
				events := tx.Events
				for _, event := range events {
					checkEventForShow(eventName, event, hasAttributeName, value, block, showAttributeName)
				}
			}
			select {
			case <-signalChan:
				return
			case <-ticker.C:
				if !disableInteractive {
					fmt.Printf("Current Block: %d\r", block)
				}
			default:
			}
		}
	}

	if blocks > 0 {
		if fromBlock <= 0 {
			fromBlock = latestHeight - blocks
		}
		utils.LavaFormatInfo("Reading Events", utils.Attribute{Key: "from", Value: fromBlock}, utils.Attribute{Key: "to", Value: fromBlock + blocks})
		readEventsFromBlock(fromBlock, fromBlock+blocks, "")
	}
	lavaChainFetcher := chainlib.NewLavaChainFetcher(ctx, clientCtx)
	latestBlock, err := lavaChainFetcher.FetchLatestBlockNum(ctx)
	if err != nil {
		return utils.LavaFormatError("failed reading latest block", err)
	}
	if shouldBreak {
		return nil
	}
	utils.LavaFormatInfo("Reading blocks Forward", utils.Attribute{Key: "current", Value: latestBlock})
	blocksToSaveChainTracker := uint64(10) // to avoid reading the same thing twice
	chainTrackerConfig := chaintracker.ChainTrackerConfig{
		BlocksToSave:      blocksToSaveChainTracker,
		AverageBlockTime:  10 * time.Second,
		ServerBlockMemory: 100 + blocksToSaveChainTracker,
		NewLatestCallback: readEventsFromBlock,
	}
	chainTracker, err := chaintracker.NewChainTracker(ctx, lavaChainFetcher, chainTrackerConfig)
	if err != nil {
		return utils.LavaFormatError("failed setting up chain tracker", err)
	}
	_ = chainTracker
	select {
	case <-ctx.Done():
		utils.LavaFormatInfo("events ctx.Done")
	case <-signalChan:
		utils.LavaFormatInfo("events signalChan")
	}
	return nil
}

func checkEventForShow(eventName string, event types.Event, hasAttributeName string, value string, block int64, showAttributeName string) {
	printEvent := func(event types.Event, showAttributeName string) string {
		attributesFilter := map[string]struct{}{}
		if showAttributeName != "" {
			attributes := strings.Split(showAttributeName, " ")
			for _, attr := range attributes {
				attributesFilter[attr] = struct{}{}
			}
		}
		passFilter := func(attr types.EventAttribute) bool {
			if len(attributesFilter) == 0 {
				return true
			}
			for attrName := range attributesFilter {
				if strings.Contains(attr.Key, attrName) {
					return true
				}
			}
			return false
		}
		st := event.Type + ": "
		sort.Slice(event.Attributes, func(i, j int) bool {
			return event.Attributes[i].Key < event.Attributes[j].Key
		})
		stmore := ""
		for _, attr := range event.Attributes {
			if passFilter(attr) {
				stmore += fmt.Sprintf("%s = %s, ", attr.Key, attr.Value)
			}
		}
		if stmore == "" {
			return ""
		}
		return st + stmore
	}
	if eventName == "" || strings.Contains(event.Type, eventName) {
		printEventTriggerValue := false
		printEventTriggerHasAttr := false
		printEventAttribute := ""
		for _, attribute := range event.Attributes {
			if hasAttributeName == "" || strings.Contains(attribute.Key, hasAttributeName) {
				printEventTriggerHasAttr = true
			}
			if value == "" || strings.Contains(attribute.Value, value) {
				printEventTriggerValue = true
			}
		}
		if printEventTriggerHasAttr && printEventTriggerValue && printEventAttribute == "" {
			printEventData := printEvent(event, showAttributeName)
			if printEventData != "" {
				utils.LavaFormatInfo("Found event", utils.Attribute{Key: "event", Value: printEventData}, utils.Attribute{Key: "height", Value: block})
			}
		}
	}
}

func CreateEventsCobraCommand() *cobra.Command {
	cmdEvents := &cobra.Command{
		Use:   `events <blocks(int)> [start_block(int)] {--value keyword | --event event_name | --from <wallet>} [--timeout duration]`,
		Short: `reads events from the current block and backwards and prints on match criteria, after it's done reads events forward`,
		Long: `reads events from the current block and backwards and prints on match criteria, after it's done reads events forward
blocks is the amount of blocks to read, when provided without a start_block will read the last X blocks going back from the current one, 0 will only read forward from now
start_blocks is an optional argument to specify the block you want to start reading events from, in case you have a specific block range you need
you must specify either: --value/--event/--from flags
--value & --event can be used at the same time, from & value conflict`,
		Example: `lavad test events 100 --event lava_relay_payment // show all events of the name lava_relay_payment from current-block - 100 and forwards
lavad test events 0 --from servicer1 // show all events from current block forwards that has my wallet address in one of their fields
lavad test events 100 5000 --value banana // show all events from 5000-5100 and current block forward that has in one of their fields the string banana
		`,
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// handle flags, pass necessary fields
			ctx := context.Background()
			networkChainId, err := cmd.Flags().GetString(flags.FlagChainID)
			if err != nil {
				return err
			}
			logLevel, err := cmd.Flags().GetString(flags.FlagLogLevel)
			if err != nil {
				utils.LavaFormatFatal("failed to read log level flag", err)
			}

			value, err := cmd.Flags().GetString(FlagValue)
			if err != nil {
				utils.LavaFormatFatal("failed to read value flag", err)
			}
			if value == "" {
				// look for a value that is from the --from flag
				from, err := cmd.Flags().GetString(flags.FlagFrom)
				if err != nil {
					utils.LavaFormatFatal("failed to read from flag", err)
				}
				if from != "" {
					keyName, err := sigs.GetKeyName(clientCtx)
					if err != nil {
						utils.LavaFormatFatal("failed getting key name from clientCtx, either provider the address in an argument or verify the --from wallet exists", err)
					}
					clientKey, err := clientCtx.Keyring.Key(keyName)
					if err != nil {
						return err
					}
					addr, err := clientKey.GetAddress()
					if err != nil {
						return err
					}
					value = addr.String()
				}
			}
			eventName, err := cmd.Flags().GetString(FlagEventName)
			if err != nil {
				utils.LavaFormatFatal("failed to read --event flag", err)
			}
			hasAttirbuteName, err := cmd.Flags().GetString(FlagHasAttributeName)
			if err != nil {
				utils.LavaFormatFatal("failed to read --attribute flag", err)
			}
			showAttirbuteName, err := cmd.Flags().GetString(FlagShowAttributeName)
			if err != nil {
				utils.LavaFormatFatal("failed to read --attribute flag", err)
			}
			blocks, err := strconv.ParseInt(args[0], 0, 64)
			if err != nil {
				utils.LavaFormatFatal("failed to parse blocks as a number", err)
			}
			if blocks < 0 {
				blocks = 0
			}

			fromBlock := int64(-1)
			if len(args) == 2 {
				fromBlock, err = strconv.ParseInt(args[1], 0, 64)
				if err != nil {
					utils.LavaFormatFatal("failed to parse blocks as a number", err)
				}
			}

			timeout, err := cmd.Flags().GetDuration(FlagTimeout)
			if err != nil {
				utils.LavaFormatFatal("failed to fetch timeout flag", err)
			}

			shouldBreak, err := cmd.Flags().GetBool(FlagBreak)
			if err != nil {
				utils.LavaFormatFatal("failed to fetch break flag", err)
			}

			disableInteractive, err := cmd.Flags().GetBool(FlagDisableInteractiveShell)
			if err != nil {
				utils.LavaFormatFatal("failed to fetch DisableInteractive flag", err)
			}
			utils.LavaFormatInfo("Events Lookup started", utils.Attribute{Key: "blocks", Value: blocks})
			utils.SetGlobalLoggingLevel(logLevel)
			clientCtx = clientCtx.WithChainID(networkChainId)
			_, err = tx.NewFactoryCLI(clientCtx, cmd.Flags())
			if err != nil {
				utils.LavaFormatFatal("failed to parse blocks as a number", err)
			}
			utils.LavaFormatInfo("lavad Binary Version: " + version.Version)
			rand.InitRandomSeed()
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			return eventsLookup(ctx, clientCtx, blocks, fromBlock, eventName, value, shouldBreak, hasAttirbuteName, showAttirbuteName, disableInteractive)
		},
	}
	flags.AddQueryFlagsToCmd(cmdEvents)
	flags.AddKeyringFlags(cmdEvents.Flags())
	cmdEvents.Flags().String(flags.FlagFrom, "", "Name or address of wallet from which to read address, and look for it in value")
	cmdEvents.Flags().Duration(FlagTimeout, 5*time.Minute, "the time to listen for events, defaults to 5m")
	cmdEvents.Flags().String(FlagValue, "", "used to show only events that has this value in one of the attributes")
	cmdEvents.Flags().Bool(FlagBreak, false, "if true will break after reading the specified amount of blocks instead of listening forward")
	cmdEvents.Flags().String(FlagEventName, "", "event name/type to look for")
	cmdEvents.Flags().String(flags.FlagChainID, app.Name, "network chain id")
	cmdEvents.Flags().String(FlagHasAttributeName, "", "only show events containing specific attribute name")
	cmdEvents.Flags().String(FlagShowAttributeName, "", "only show a specific attribute name, and no other attributes")
	cmdEvents.Flags().Bool(FlagDisableInteractiveShell, false, "a flag to disable the shell printing interactive prints, used when scripting the command")
	return cmdEvents
}

func CreateTxCounterCobraCommand() *cobra.Command {
	cmdTxCounter := &cobra.Command{
		Use:     `txcounter  [number_of_days_to_count(int)] [average_block_time_in_seconds(int)]`,
		Short:   `tx counter  [number_of_days_to_count(int)] [average_block_time_in_seconds(int)]`,
		Long:    `tx counter  [number_of_days_to_count(int)] [average_block_time_in_seconds(int)]`,
		Example: `lavad test txcounter 1 15 -- will count 1 day worth of blocks where each block is 15 seconds`,
		Args:    cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// handle flags, pass necessary fields
			ctx := context.Background()
			networkChainId, err := cmd.Flags().GetString(flags.FlagChainID)
			if err != nil {
				return err
			}
			logLevel, err := cmd.Flags().GetString(flags.FlagLogLevel)
			if err != nil {
				utils.LavaFormatFatal("failed to read log level flag", err)
			}

			numberOfDays := int64(-1)
			if len(args) == 2 {
				numberOfDays, err = strconv.ParseInt(args[0], 0, 64)
				if err != nil {
					utils.LavaFormatFatal("failed to parse blocks as a number", err)
				}
			}

			blockTime, err := strconv.ParseInt(args[1], 0, 64)
			if err != nil {
				utils.LavaFormatFatal("failed to parse blocks as a number", err)
			}
			if blockTime < 0 {
				blockTime = 0
			}

			utils.LavaFormatInfo("Events Lookup started", utils.Attribute{Key: "blocks", Value: blockTime})
			utils.SetGlobalLoggingLevel(logLevel)
			clientCtx = clientCtx.WithChainID(networkChainId)
			_, err = tx.NewFactoryCLI(clientCtx, cmd.Flags())
			if err != nil {
				utils.LavaFormatFatal("failed to parse blocks as a number", err)
			}
			utils.LavaFormatInfo("lavad Binary Version: " + version.Version)
			rand.InitRandomSeed()
			return countTransactionsPerDay(ctx, clientCtx, blockTime, numberOfDays)
		},
	}
	flags.AddQueryFlagsToCmd(cmdTxCounter)
	flags.AddKeyringFlags(cmdTxCounter.Flags())
	cmdTxCounter.Flags().String(flags.FlagFrom, "", "Name or address of wallet from which to read address, and look for it in value")
	cmdTxCounter.Flags().Duration(FlagTimeout, 5*time.Minute, "the time to listen for events, defaults to 5m")
	cmdTxCounter.Flags().String(flags.FlagChainID, app.Name, "network chain id")
	return cmdTxCounter
}

func countTransactionsPerDay(ctx context.Context, clientCtx client.Context, blockTime, numberOfDays int64) error {
	resultStatus, err := clientCtx.Client.Status(ctx)
	if err != nil {
		return err
	}
	latestHeight := resultStatus.SyncInfo.LatestBlockHeight
	// 1 block for blockTime (lets say 15 seconds)
	// number of seconds in a day: 24 * 60 * 60
	numberOfSecondsInADay := int64(24 * 60 * 60)
	numberOfBlocksInADay := numberOfSecondsInADay / blockTime
	utils.LavaFormatInfo("Starting counter",
		utils.LogAttr("latest_block", latestHeight),
		utils.LogAttr("numberOfSecondsInADay", numberOfSecondsInADay),
		utils.LogAttr("numberOfBlocksInADay", numberOfBlocksInADay),
		utils.LogAttr("starting_block", latestHeight-numberOfBlocksInADay),
	)

	tmClient, err := updaters.TryIntoTendermintRPC(clientCtx.Client)
	if err != nil {
		utils.LavaFormatFatal("invalid blockResults provider", err)
	}
	// i is days
	// j are blocks in that day
	// starting from current day and going backwards
	totalTxPerDay := map[int64]int{}
	for i := int64(1); i <= numberOfDays; i++ {
		utils.LavaFormatInfo("Parsing day", utils.LogAttr("Day", i), utils.LogAttr("starting block", latestHeight-(numberOfBlocksInADay*i)), utils.LogAttr("ending block", latestHeight-(numberOfBlocksInADay*(i-1))))
		for j := latestHeight - (numberOfBlocksInADay * i); j < latestHeight-(numberOfBlocksInADay*(i-1)); j++ {
			ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
			blockResults, err := tmClient.BlockResults(ctxWithTimeout, &j)
			cancel()
			if err != nil {
				utils.LavaFormatError("invalid blockResults status", err)
				continue
			}
			transactionResults := blockResults.TxsResults
			utils.LavaFormatInfo("Number of tx for block", utils.LogAttr("block_number", j), utils.LogAttr("number_of_tx", len(transactionResults)))
			if _, ok := totalTxPerDay[i]; ok {
				newLength := totalTxPerDay[i] + len(transactionResults)
				totalTxPerDay[i] = newLength
			} else {
				totalTxPerDay[i] = len(transactionResults)
			}
		}
	}
	utils.LavaFormatInfo("transactions per day results", utils.LogAttr("totalTxPerDay", totalTxPerDay))

	// Create a map to hold the JSON data
	jsonData := make(map[string]int)
	for key, value := range totalTxPerDay {
		// Calculate the date for each key
		date := time.Now().AddDate(0, 0, -int(key)+1).Format("2006-01-02")
		dateKey := fmt.Sprintf("date_%s", date)
		jsonData[dateKey] = value
	}

	// Convert the JSON data to JSON format
	jsonBytes, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	// Write JSON data to a file
	fileName := "dates.json"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonBytes)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	utils.LavaFormatInfo("JSON data has been written to:" + fileName)
	return nil

	// http://testnet2-rpc.lavapro.xyz/
}
