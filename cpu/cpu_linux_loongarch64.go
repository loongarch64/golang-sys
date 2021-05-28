// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loongarch64
// +build linux
// +build loongarch64

package cpu

// HWCAP/HWCAP2 bits. These are exposed by the Linux kernel.
const (
	/* HWCAP flags */
	hwcap_LOONGARCH_CPUCFG = 1 << 0
	hwcap_LOONGARCH_FPU    = 1 << 1
	hwcap_LOONGARCH_LSX    = 1 << 2
	hwcap_LOONGARCH_LASX   = 1 << 3
	hwcap_LOONGARCH_LBT    = 1 << 4
	hwcap_LOONGARCH_LVZ    = 1 << 5
	hwcap_LOONGARCH_AES    = 1 << 6
	hwcap_LOONGARCH_CRC32  = 1 << 7
	hwcap_LOONGARCH_SHA1   = 1 << 8
	hwcap_LOONGARCH_SHA2   = 1 << 9
	hwcap_LOONGARCH_SHA3   = 1 << 10
)

func doinit() {
	LOONGARCH64.HasCPUCFG = isSet(hwCap, hwcap_LOONGARCH_CPUCFG)
	LOONGARCH64.HasFPU = isSet(hwCap, hwcap_LOONGARCH_FPU)
	LOONGARCH64.HasLSX = isSet(hwCap, hwcap_LOONGARCH_LSX)
	LOONGARCH64.HasLASX = isSet(hwCap, hwcap_LOONGARCH_LASX)
	LOONGARCH64.HasLBT = isSet(hwCap, hwcap_LOONGARCH_LBT)
	LOONGARCH64.HasLVZ = isSet(hwCap, hwcap_LOONGARCH_LVZ)
	LOONGARCH64.HasAES = isSet(hwCap, hwcap_LOONGARCH_AES)
	LOONGARCH64.HasCRC32 = isSet(hwCap, hwcap_LOONGARCH_CRC32)
	LOONGARCH64.HasSHA1 = isSet(hwCap, hwcap_LOONGARCH_SHA1)
	LOONGARCH64.HasSHA2 = isSet(hwCap, hwcap_LOONGARCH_SHA2)
	LOONGARCH64.HasSHA3 = isSet(hwCap, hwcap_LOONGARCH_SHA3)
}

func isSet(hwc uint, value uint) bool {
	return hwc&value != 0
}
