package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

// Config
const (
	ServerURL      = "http://127.0.0.1:5000/commandz"
	PollInterval   = 10 * time.Second
	StatsInterval  = 5 * time.Second
)

func main() {
	fmt.Println("Starting Nagnos Go Agent (v2)...")

	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine for Dynamic Analysis (fetching commands from backend)
	go func() {
		defer wg.Done()
		fetchCommandsLoop()
	}()

	// Goroutine for Static Analysis (collecting local system stats)
	go func() {
		defer wg.Done()
		collectStatsLoop()
	}()

	wg.Wait()
}

// fetchCommandsLoop runs indefinitely, polling the server for commands.
func fetchCommandsLoop() {
	ticker := time.NewTicker(PollInterval)
	defer ticker.Stop()

	// Run immediately once
	if cmds, err := fetchCommands(); err == nil {
		fmt.Printf("[Dynamic Analysis] Fetched %d commands from server\n", len(cmds))
	} else {
		fmt.Printf("[Dynamic Analysis] Error fetching commands: %v\n", err)
	}

	for {
		select {
		case <-ticker.C:
			commands, err := fetchCommands()
			if err != nil {
				fmt.Printf("[Dynamic Analysis] Error fetching commands: %v\n", err)
			} else {
				fmt.Printf("[Dynamic Analysis] Fetched %d commands from server\n", len(commands))
			}
		}
	}
}

// fetchCommands performs the HTTP request to the backend.
func fetchCommands() ([]string, error) {
	resp, err := http.Get(ServerURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Expected JSON: {"commands": [["ls"], ["cd"], ...]}
	type CommandResponse struct {
		Commands [][]string `json:"commands"`
	}

	var cmdResp CommandResponse
	if err := json.Unmarshal(body, &cmdResp); err != nil {
		return nil, err
	}

	var flatCommands []string
	for _, cmdPair := range cmdResp.Commands {
		if len(cmdPair) > 0 {
			flatCommands = append(flatCommands, cmdPair[0])
		}
	}
	return flatCommands, nil
}

// collectStatsLoop runs indefinitely, gathering and printing system statistics.
func collectStatsLoop() {
	ticker := time.NewTicker(StatsInterval)
	defer ticker.Stop()

	// Run immediately once
	printStats()

	for {
		select {
		case <-ticker.C:
			printStats()
		}
	}
}

// printStats gathers and prints the statistics.
func printStats() {
	stats := getSystemStats()
	fmt.Printf("[Static Analysis] OS: %s | CPU Cores: %d | Goroutines: %d | MemAlloc: %d bytes | LoadAvg: %s | Disk Free: %d GB\n",
		stats.GOOS, stats.NumCPU, stats.NumGoroutine, stats.MemAlloc, stats.LoadAvg, stats.DiskFreeGB)
}

type SystemStats struct {
	NumCPU       int
	NumGoroutine int
	GOOS         string
	MemAlloc     uint64
	LoadAvg      string
	DiskFreeGB   uint64
}

// getSystemStats collects various system metrics.
func getSystemStats() SystemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	loadAvg := getLoadAvg()
	diskFree := getDiskFree("/")

	return SystemStats{
		NumCPU:       runtime.NumCPU(),
		NumGoroutine: runtime.NumGoroutine(),
		GOOS:         runtime.GOOS,
		MemAlloc:     m.Alloc,
		LoadAvg:      loadAvg,
		DiskFreeGB:   diskFree,
	}
}

// getLoadAvg reads /proc/loadavg (Unix/Linux specific)
func getLoadAvg() string {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return "N/A"
	}
	parts := strings.Fields(string(data))
	if len(parts) >= 3 {
		return strings.Join(parts[:3], " ")
	}
	return string(data)
}

// getDiskFree uses syscall.Statfs to get free disk space (Unix/Linux specific)
func getDiskFree(path string) uint64 {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return 0
	}
	// Available blocks * size per block = available space in bytes
	// Note: Bsize is int64 on some archs, so we cast to uint64 to match Bavail
	freeBytes := stat.Bavail * uint64(stat.Bsize)
	return freeBytes / (1024 * 1024 * 1024)
}
