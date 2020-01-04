package nes

import "time"

const (
    DefaultFPSNTSC float64 = 60.0988
    DefaultFPSPAL  float64 = 50.0070
)

type FPS struct {
    enabled bool
    frames  float64
    rate    float64
    ticks   uint64
}

func NewFPS(rate float64) *FPS {
    fps := &FPS{}
    fps.SetRate(rate)
    return fps
}

func (fps *FPS) Enable() {
    fps.enabled = true
}

func (fps *FPS) Disable() {
    fps.enabled = false
}

func (fps *FPS) Delay() {
    fps.frames++
    current := uint64(time.Now().UnixNano()) / 1e6
    target := fps.ticks + uint64(fps.frames * fps.rate)

    if fps.enabled && current <= target {
        time.Sleep(time.Duration((target - current) * 1e6)
    } else {
        fps.frames = 0.0
        fps.ticks = uint64(time.Now().UnixNano()) / 1e6
    }
}

