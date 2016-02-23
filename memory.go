package gob

import (
	"errors"
	_ "fmt"
	"io/ioutil"
)

type MemoryIO struct {
	ROM      []uint8
	WRAM     [8192]uint8
	CRAM     []uint8
	ROMBank  uint8
	CRAMBank uint8
}

func (m *MemoryIO) loadROM(filename string) error {
	var err error

	m.ROM, err = ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if m.GetUint8(0x147) != 0 {
		// MBC setup, error for now
		return errors.New("MBC not supported")
	}

	return nil
}

func (m *MemoryIO) GetUint8(addr uint16) uint8 {
	if addr >= 0x0000 && addr <= 0x3FFF {
		// unbanked ROM (the first 16kb)
		return m.ROM[addr]
	} else if addr >= 0x4000 && addr <= 0x7FFF {
		// banked ROM (16kb+)
	} else if addr >= 0x8000 && addr <= 0x9FFF {
		// VRAM
	} else if addr >= 0xA000 && addr <= 0xBFFF {
		// banked cartridge RAM
	} else if addr >= 0xC000 && addr <= 0xCFFF {
		// Work RAM (Bank 0)
	} else if addr >= 0xD000 && addr <= 0xDFFF {
		// Work RAM (Bank 1)
	} else if addr >= 0xE000 && addr <= 0xFDFF {
		// C000-DDFF mirror
	} else if addr >= 0xFE00 && addr <= 0xFE9F {
		// OAM (sprite attributes)
	} else if addr >= 0xFEA0 && addr <= 0xFEFF {
		// unusable
	} else if addr >= 0xFF00 && addr <= 0xFF7F {
		// I/O Ports
	} else if addr >= 0xFF80 && addr <= 0xFFFE {
		// hRAM
	} else if addr == 0xFFFF {
		// Interrupt Enable Register
	}

	return 0xFF
}

func (m *MemoryIO) GetUint16(addr uint16) uint16 {
	// the GB is little-endian. make sure to swap the bytes
	return (uint16(m.GetUint8(addr+1)) << 8) | uint16(m.GetUint8(addr))
}
