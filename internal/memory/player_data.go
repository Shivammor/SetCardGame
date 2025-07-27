package memory

import (
    "log"
    "sync"
)

type PlayerData struct {
    Name       string
    Avatar     int    // Index of selected avatar (0-3)
    RoomKey    string
    LastUsed   int64  // Unix timestamp
}

type MemoryStore struct {
    mu         sync.RWMutex
    playerData PlayerData
}

var globalStore = &MemoryStore{}

func GetPlayerData() PlayerData {
    globalStore.mu.RLock()
    defer globalStore.mu.RUnlock()
    return globalStore.playerData
}

func SetPlayerData(data PlayerData) {
    globalStore.mu.Lock()
    defer globalStore.mu.Unlock()
    globalStore.playerData = data
    log.Printf("💾 Saved player data: Name=%s, Avatar=%d, RoomKey=%s", 
        data.Name, data.Avatar, data.RoomKey)
}

func SetRoomKey(roomKey string) {
    globalStore.mu.Lock()
    defer globalStore.mu.Unlock()
    globalStore.playerData.RoomKey = roomKey
    log.Printf("🔑 Saved room key: %s", roomKey)
}

func GetRoomKey() string {
    globalStore.mu.RLock()
    defer globalStore.mu.RUnlock()
    return globalStore.playerData.RoomKey
}

func SetPlayerName(name string) {
    globalStore.mu.Lock()
    defer globalStore.mu.Unlock()
    globalStore.playerData.Name = name
    log.Printf("👤 Saved player name: %s", name)
}

func GetPlayerName() string {
    globalStore.mu.RLock()
    defer globalStore.mu.RUnlock()
    return globalStore.playerData.Name
}

func SetSelectedAvatar(avatar int) {
    globalStore.mu.Lock()
    defer globalStore.mu.Unlock()
    globalStore.playerData.Avatar = avatar
    log.Printf("🎭 Saved selected avatar: %d", avatar)
}

func GetSelectedAvatar() int {
    globalStore.mu.RLock()
    defer globalStore.mu.RUnlock()
    return globalStore.playerData.Avatar
}

