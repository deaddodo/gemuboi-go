package memory

import (
	"errors"
	"io/ioutil"
)

// IO provides an emulated memory interface
type IO struct {
	ROM      []uint8
	BIOS     []uint8
	WRAM     [8192]uint8
	VRAM     [8192]uint8
	CRAM     []uint8
	ROMBank  uint8
	CRAMBank uint8
	BIOSMode bool
}

// Init dumps the provided ROM into memory, sets some internal values (MBC,
// RAM bank, ROM bank, etc). We want this data independent of the CPU in case of
// soft reset
func (m *IO) Init(biosFile string, romFile string) error {
	var err error

	m.BIOS, err = ioutil.ReadFile(biosFile)
	if err != nil {
		return err
	} else if len(m.BIOS) != 256 {
		return errors.New("ROM file incorrect")
	}

	m.ROM, err = ioutil.ReadFile(romFile)
	if err != nil {
		return err
	}

	if m.GetUint8(0x147) != 0 {
		// MBC setup, error for now
		return errors.New("MBC not supported")
	}

	m.ROMBank = 1
	m.BIOSMode = true

	return nil
}

func (m *IO) readIO(addr uint8) uint8 {
	switch addr {
	case 0x00:
		// joypad status
		fallthrough
	case 0x01, 0x02:
		// link cable
		fallthrough
	case 0x04, 0x05, 0x06, 0x07:
		// timer/divider
		fallthrough
	case 0x0f:
		// interrupt flag
		fallthrough
	default:
		return 0xFF
	}
}

// GetUint8 pulls a standard byte value from memory, by address
func (m *IO) GetUint8(addr uint16) uint8 {
	switch {
	case addr <= 0x3FFF:
		if m.BIOSMode && (addr <= 0xFF) {
			return m.BIOS[addr]
		}
		return m.ROM[addr]
	case addr >= 0x4000 && addr <= 0x7FFF:
		return m.ROM[(uint32(addr)-0x4000)+(0x4000*uint32(m.ROMBank))]
	case addr >= 0x8000 && addr <= 0x9FFF:
		// TODO: GBC: Make Bankable (VRAM / Read)
		return m.VRAM[addr-0x8000]
	case addr >= 0xA000 && addr <= 0xBFFF:
		// TODO: banked cartridge RAM (Read)
		fallthrough
	case addr >= 0xC000 && addr <= 0xCFFF:
		return m.WRAM[addr-0xC000]
	case addr >= 0xD000 && addr <= 0xDFFF:
		// TODO: GBC: Make Bankable (Work RAM / Read)
		return m.WRAM[addr-0xD000]
	case addr >= 0xE000 && addr <= 0xFDFF:
		return m.GetUint8(addr - 0x2000)
	case addr >= 0xFE00 && addr <= 0xFE9F:
		// TODO: OAM (sprite attributes)
		fallthrough
	case addr >= 0xFEA0 && addr <= 0xFEFF:
		// TODO: unusable (read)
		fallthrough
	case addr >= 0xFF00 && addr <= 0xFF7F:
		return m.readIO((uint8(addr - 0xFF00)))
	case addr >= 0xFF80 && addr <= 0xFFFE:
		// TODO: hRAM (Read)
		fallthrough
	default:
		// TODO: 0xFFFF: Interrupt Enable Register (Read)
		return 0xFF
	}
}

// GetUint16 pulls a 16-bit value from memory. Since this is little-endian,
// swap the values to represent their logical value (50h 01h -> 150h)
func (m *IO) GetUint16(addr uint16) uint16 {
	return (uint16(m.GetUint8(addr+1)) << 8) | uint16(m.GetUint8(addr))
}

// SetUint8 set a memory address to a byte value
func (m *IO) SetUint8(addr uint16, value uint8) {
	switch {
	case addr <= 0x7FFF:
		// TODO: MBC control occurs in this range
	case addr >= 0x8000 && addr <= 0x9FFF:
		// TODO: GBC: Make Bankable (VRAM / Write)
		m.VRAM[addr-0x8000] = value
	case addr >= 0xA000 && addr <= 0xBFFF:
		// TODO: banked cart RAM (possible external device / Write)
		fallthrough
	case addr >= 0xC000 && addr <= 0xCFFF:
		m.WRAM[addr-0xC000] = value
	case addr >= 0xD000 && addr <= 0xDFFF:
		// TODO: GBC: Make Bankable (Work RAM / Write)
		m.WRAM[addr-0xD000] = value
	case addr >= 0xE000 && addr <= 0xFDFF:
		// TODO: How does ECHO RAM handle writes?
		//m.GetUint8(addr - 0x2000)
	case addr >= 0xFE00 && addr <= 0xFE9F:
		// TODO: OAM (sprite attributes)
		fallthrough
	case addr >= 0xFEA0 && addr <= 0xFEFF:
		// TODO: How do writes to the unusable area work?
		fallthrough
	case addr >= 0xFF00 && addr <= 0xFF7F:
		// TODO: I/O Ports (write)
		// m.writeIO((uint8(addr - 0xFF00)))
	case addr >= 0xFF80 && addr <= 0xFFFE:
		// TODO: hRAM (Write)
		fallthrough
	default:
		// TODO: 0xFFFF: Interrupt Enable Register (Write)
	}
}

// SetUint16 set a memory address to a double-byte value
func (m *IO) SetUint16(addr uint16, value uint16) {
	m.SetUint8(addr+1, uint8(value))
	m.SetUint8(addr, uint8(value>>8))
}
