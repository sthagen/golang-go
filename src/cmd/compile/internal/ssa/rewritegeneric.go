// Code generated from gen/generic.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "math"
import "cmd/compile/internal/types"

func rewriteValuegeneric(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		return rewriteValuegeneric_OpAdd16(v)
	case OpAdd32:
		return rewriteValuegeneric_OpAdd32(v)
	case OpAdd32F:
		return rewriteValuegeneric_OpAdd32F(v)
	case OpAdd64:
		return rewriteValuegeneric_OpAdd64(v)
	case OpAdd64F:
		return rewriteValuegeneric_OpAdd64F(v)
	case OpAdd8:
		return rewriteValuegeneric_OpAdd8(v)
	case OpAddPtr:
		return rewriteValuegeneric_OpAddPtr(v)
	case OpAnd16:
		return rewriteValuegeneric_OpAnd16(v)
	case OpAnd32:
		return rewriteValuegeneric_OpAnd32(v)
	case OpAnd64:
		return rewriteValuegeneric_OpAnd64(v)
	case OpAnd8:
		return rewriteValuegeneric_OpAnd8(v)
	case OpArraySelect:
		return rewriteValuegeneric_OpArraySelect(v)
	case OpCom16:
		return rewriteValuegeneric_OpCom16(v)
	case OpCom32:
		return rewriteValuegeneric_OpCom32(v)
	case OpCom64:
		return rewriteValuegeneric_OpCom64(v)
	case OpCom8:
		return rewriteValuegeneric_OpCom8(v)
	case OpConstInterface:
		return rewriteValuegeneric_OpConstInterface(v)
	case OpConstSlice:
		return rewriteValuegeneric_OpConstSlice(v)
	case OpConstString:
		return rewriteValuegeneric_OpConstString(v)
	case OpConvert:
		return rewriteValuegeneric_OpConvert(v)
	case OpCvt32Fto32:
		return rewriteValuegeneric_OpCvt32Fto32(v)
	case OpCvt32Fto64:
		return rewriteValuegeneric_OpCvt32Fto64(v)
	case OpCvt32Fto64F:
		return rewriteValuegeneric_OpCvt32Fto64F(v)
	case OpCvt32to32F:
		return rewriteValuegeneric_OpCvt32to32F(v)
	case OpCvt32to64F:
		return rewriteValuegeneric_OpCvt32to64F(v)
	case OpCvt64Fto32:
		return rewriteValuegeneric_OpCvt64Fto32(v)
	case OpCvt64Fto32F:
		return rewriteValuegeneric_OpCvt64Fto32F(v)
	case OpCvt64Fto64:
		return rewriteValuegeneric_OpCvt64Fto64(v)
	case OpCvt64to32F:
		return rewriteValuegeneric_OpCvt64to32F(v)
	case OpCvt64to64F:
		return rewriteValuegeneric_OpCvt64to64F(v)
	case OpCvtBoolToUint8:
		return rewriteValuegeneric_OpCvtBoolToUint8(v)
	case OpDiv16:
		return rewriteValuegeneric_OpDiv16(v)
	case OpDiv16u:
		return rewriteValuegeneric_OpDiv16u(v)
	case OpDiv32:
		return rewriteValuegeneric_OpDiv32(v)
	case OpDiv32F:
		return rewriteValuegeneric_OpDiv32F(v)
	case OpDiv32u:
		return rewriteValuegeneric_OpDiv32u(v)
	case OpDiv64:
		return rewriteValuegeneric_OpDiv64(v)
	case OpDiv64F:
		return rewriteValuegeneric_OpDiv64F(v)
	case OpDiv64u:
		return rewriteValuegeneric_OpDiv64u(v)
	case OpDiv8:
		return rewriteValuegeneric_OpDiv8(v)
	case OpDiv8u:
		return rewriteValuegeneric_OpDiv8u(v)
	case OpEq16:
		return rewriteValuegeneric_OpEq16(v)
	case OpEq32:
		return rewriteValuegeneric_OpEq32(v)
	case OpEq32F:
		return rewriteValuegeneric_OpEq32F(v)
	case OpEq64:
		return rewriteValuegeneric_OpEq64(v)
	case OpEq64F:
		return rewriteValuegeneric_OpEq64F(v)
	case OpEq8:
		return rewriteValuegeneric_OpEq8(v)
	case OpEqB:
		return rewriteValuegeneric_OpEqB(v)
	case OpEqInter:
		return rewriteValuegeneric_OpEqInter(v)
	case OpEqPtr:
		return rewriteValuegeneric_OpEqPtr(v)
	case OpEqSlice:
		return rewriteValuegeneric_OpEqSlice(v)
	case OpGeq32F:
		return rewriteValuegeneric_OpGeq32F(v)
	case OpGeq64F:
		return rewriteValuegeneric_OpGeq64F(v)
	case OpGreater32F:
		return rewriteValuegeneric_OpGreater32F(v)
	case OpGreater64F:
		return rewriteValuegeneric_OpGreater64F(v)
	case OpIMake:
		return rewriteValuegeneric_OpIMake(v)
	case OpInterCall:
		return rewriteValuegeneric_OpInterCall(v)
	case OpIsInBounds:
		return rewriteValuegeneric_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValuegeneric_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValuegeneric_OpIsSliceInBounds(v)
	case OpLeq16:
		return rewriteValuegeneric_OpLeq16(v)
	case OpLeq16U:
		return rewriteValuegeneric_OpLeq16U(v)
	case OpLeq32:
		return rewriteValuegeneric_OpLeq32(v)
	case OpLeq32F:
		return rewriteValuegeneric_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValuegeneric_OpLeq32U(v)
	case OpLeq64:
		return rewriteValuegeneric_OpLeq64(v)
	case OpLeq64F:
		return rewriteValuegeneric_OpLeq64F(v)
	case OpLeq64U:
		return rewriteValuegeneric_OpLeq64U(v)
	case OpLeq8:
		return rewriteValuegeneric_OpLeq8(v)
	case OpLeq8U:
		return rewriteValuegeneric_OpLeq8U(v)
	case OpLess16:
		return rewriteValuegeneric_OpLess16(v)
	case OpLess16U:
		return rewriteValuegeneric_OpLess16U(v)
	case OpLess32:
		return rewriteValuegeneric_OpLess32(v)
	case OpLess32F:
		return rewriteValuegeneric_OpLess32F(v)
	case OpLess32U:
		return rewriteValuegeneric_OpLess32U(v)
	case OpLess64:
		return rewriteValuegeneric_OpLess64(v)
	case OpLess64F:
		return rewriteValuegeneric_OpLess64F(v)
	case OpLess64U:
		return rewriteValuegeneric_OpLess64U(v)
	case OpLess8:
		return rewriteValuegeneric_OpLess8(v)
	case OpLess8U:
		return rewriteValuegeneric_OpLess8U(v)
	case OpLoad:
		return rewriteValuegeneric_OpLoad(v)
	case OpLsh16x16:
		return rewriteValuegeneric_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValuegeneric_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValuegeneric_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValuegeneric_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValuegeneric_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValuegeneric_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValuegeneric_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValuegeneric_OpLsh32x8(v)
	case OpLsh64x16:
		return rewriteValuegeneric_OpLsh64x16(v)
	case OpLsh64x32:
		return rewriteValuegeneric_OpLsh64x32(v)
	case OpLsh64x64:
		return rewriteValuegeneric_OpLsh64x64(v)
	case OpLsh64x8:
		return rewriteValuegeneric_OpLsh64x8(v)
	case OpLsh8x16:
		return rewriteValuegeneric_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValuegeneric_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValuegeneric_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValuegeneric_OpLsh8x8(v)
	case OpMod16:
		return rewriteValuegeneric_OpMod16(v)
	case OpMod16u:
		return rewriteValuegeneric_OpMod16u(v)
	case OpMod32:
		return rewriteValuegeneric_OpMod32(v)
	case OpMod32u:
		return rewriteValuegeneric_OpMod32u(v)
	case OpMod64:
		return rewriteValuegeneric_OpMod64(v)
	case OpMod64u:
		return rewriteValuegeneric_OpMod64u(v)
	case OpMod8:
		return rewriteValuegeneric_OpMod8(v)
	case OpMod8u:
		return rewriteValuegeneric_OpMod8u(v)
	case OpMove:
		return rewriteValuegeneric_OpMove(v)
	case OpMul16:
		return rewriteValuegeneric_OpMul16(v)
	case OpMul32:
		return rewriteValuegeneric_OpMul32(v)
	case OpMul32F:
		return rewriteValuegeneric_OpMul32F(v)
	case OpMul64:
		return rewriteValuegeneric_OpMul64(v)
	case OpMul64F:
		return rewriteValuegeneric_OpMul64F(v)
	case OpMul8:
		return rewriteValuegeneric_OpMul8(v)
	case OpNeg16:
		return rewriteValuegeneric_OpNeg16(v)
	case OpNeg32:
		return rewriteValuegeneric_OpNeg32(v)
	case OpNeg32F:
		return rewriteValuegeneric_OpNeg32F(v)
	case OpNeg64:
		return rewriteValuegeneric_OpNeg64(v)
	case OpNeg64F:
		return rewriteValuegeneric_OpNeg64F(v)
	case OpNeg8:
		return rewriteValuegeneric_OpNeg8(v)
	case OpNeq16:
		return rewriteValuegeneric_OpNeq16(v)
	case OpNeq32:
		return rewriteValuegeneric_OpNeq32(v)
	case OpNeq32F:
		return rewriteValuegeneric_OpNeq32F(v)
	case OpNeq64:
		return rewriteValuegeneric_OpNeq64(v)
	case OpNeq64F:
		return rewriteValuegeneric_OpNeq64F(v)
	case OpNeq8:
		return rewriteValuegeneric_OpNeq8(v)
	case OpNeqB:
		return rewriteValuegeneric_OpNeqB(v)
	case OpNeqInter:
		return rewriteValuegeneric_OpNeqInter(v)
	case OpNeqPtr:
		return rewriteValuegeneric_OpNeqPtr(v)
	case OpNeqSlice:
		return rewriteValuegeneric_OpNeqSlice(v)
	case OpNilCheck:
		return rewriteValuegeneric_OpNilCheck(v)
	case OpNot:
		return rewriteValuegeneric_OpNot(v)
	case OpOffPtr:
		return rewriteValuegeneric_OpOffPtr(v)
	case OpOr16:
		return rewriteValuegeneric_OpOr16(v)
	case OpOr32:
		return rewriteValuegeneric_OpOr32(v)
	case OpOr64:
		return rewriteValuegeneric_OpOr64(v)
	case OpOr8:
		return rewriteValuegeneric_OpOr8(v)
	case OpPhi:
		return rewriteValuegeneric_OpPhi(v)
	case OpPtrIndex:
		return rewriteValuegeneric_OpPtrIndex(v)
	case OpRotateLeft16:
		return rewriteValuegeneric_OpRotateLeft16(v)
	case OpRotateLeft32:
		return rewriteValuegeneric_OpRotateLeft32(v)
	case OpRotateLeft64:
		return rewriteValuegeneric_OpRotateLeft64(v)
	case OpRotateLeft8:
		return rewriteValuegeneric_OpRotateLeft8(v)
	case OpRound32F:
		return rewriteValuegeneric_OpRound32F(v)
	case OpRound64F:
		return rewriteValuegeneric_OpRound64F(v)
	case OpRsh16Ux16:
		return rewriteValuegeneric_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValuegeneric_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValuegeneric_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValuegeneric_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValuegeneric_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValuegeneric_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValuegeneric_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValuegeneric_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValuegeneric_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValuegeneric_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValuegeneric_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValuegeneric_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValuegeneric_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValuegeneric_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValuegeneric_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValuegeneric_OpRsh32x8(v)
	case OpRsh64Ux16:
		return rewriteValuegeneric_OpRsh64Ux16(v)
	case OpRsh64Ux32:
		return rewriteValuegeneric_OpRsh64Ux32(v)
	case OpRsh64Ux64:
		return rewriteValuegeneric_OpRsh64Ux64(v)
	case OpRsh64Ux8:
		return rewriteValuegeneric_OpRsh64Ux8(v)
	case OpRsh64x16:
		return rewriteValuegeneric_OpRsh64x16(v)
	case OpRsh64x32:
		return rewriteValuegeneric_OpRsh64x32(v)
	case OpRsh64x64:
		return rewriteValuegeneric_OpRsh64x64(v)
	case OpRsh64x8:
		return rewriteValuegeneric_OpRsh64x8(v)
	case OpRsh8Ux16:
		return rewriteValuegeneric_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValuegeneric_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValuegeneric_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValuegeneric_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValuegeneric_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValuegeneric_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValuegeneric_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValuegeneric_OpRsh8x8(v)
	case OpSelect0:
		return rewriteValuegeneric_OpSelect0(v)
	case OpSelect1:
		return rewriteValuegeneric_OpSelect1(v)
	case OpSignExt16to32:
		return rewriteValuegeneric_OpSignExt16to32(v)
	case OpSignExt16to64:
		return rewriteValuegeneric_OpSignExt16to64(v)
	case OpSignExt32to64:
		return rewriteValuegeneric_OpSignExt32to64(v)
	case OpSignExt8to16:
		return rewriteValuegeneric_OpSignExt8to16(v)
	case OpSignExt8to32:
		return rewriteValuegeneric_OpSignExt8to32(v)
	case OpSignExt8to64:
		return rewriteValuegeneric_OpSignExt8to64(v)
	case OpSliceCap:
		return rewriteValuegeneric_OpSliceCap(v)
	case OpSliceLen:
		return rewriteValuegeneric_OpSliceLen(v)
	case OpSlicePtr:
		return rewriteValuegeneric_OpSlicePtr(v)
	case OpSlicemask:
		return rewriteValuegeneric_OpSlicemask(v)
	case OpSqrt:
		return rewriteValuegeneric_OpSqrt(v)
	case OpStaticCall:
		return rewriteValuegeneric_OpStaticCall(v)
	case OpStore:
		return rewriteValuegeneric_OpStore(v)
	case OpStringLen:
		return rewriteValuegeneric_OpStringLen(v)
	case OpStringPtr:
		return rewriteValuegeneric_OpStringPtr(v)
	case OpStructSelect:
		return rewriteValuegeneric_OpStructSelect(v)
	case OpSub16:
		return rewriteValuegeneric_OpSub16(v)
	case OpSub32:
		return rewriteValuegeneric_OpSub32(v)
	case OpSub32F:
		return rewriteValuegeneric_OpSub32F(v)
	case OpSub64:
		return rewriteValuegeneric_OpSub64(v)
	case OpSub64F:
		return rewriteValuegeneric_OpSub64F(v)
	case OpSub8:
		return rewriteValuegeneric_OpSub8(v)
	case OpTrunc16to8:
		return rewriteValuegeneric_OpTrunc16to8(v)
	case OpTrunc32to16:
		return rewriteValuegeneric_OpTrunc32to16(v)
	case OpTrunc32to8:
		return rewriteValuegeneric_OpTrunc32to8(v)
	case OpTrunc64to16:
		return rewriteValuegeneric_OpTrunc64to16(v)
	case OpTrunc64to32:
		return rewriteValuegeneric_OpTrunc64to32(v)
	case OpTrunc64to8:
		return rewriteValuegeneric_OpTrunc64to8(v)
	case OpXor16:
		return rewriteValuegeneric_OpXor16(v)
	case OpXor32:
		return rewriteValuegeneric_OpXor32(v)
	case OpXor64:
		return rewriteValuegeneric_OpXor64(v)
	case OpXor8:
		return rewriteValuegeneric_OpXor8(v)
	case OpZero:
		return rewriteValuegeneric_OpZero(v)
	case OpZeroExt16to32:
		return rewriteValuegeneric_OpZeroExt16to32(v)
	case OpZeroExt16to64:
		return rewriteValuegeneric_OpZeroExt16to64(v)
	case OpZeroExt32to64:
		return rewriteValuegeneric_OpZeroExt32to64(v)
	case OpZeroExt8to16:
		return rewriteValuegeneric_OpZeroExt8to16(v)
	case OpZeroExt8to32:
		return rewriteValuegeneric_OpZeroExt8to32(v)
	case OpZeroExt8to64:
		return rewriteValuegeneric_OpZeroExt8to64(v)
	}
	return false
}
func rewriteValuegeneric_OpAdd16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Add16 (Const16 [c]) (Const16 [d]))
	// result: (Const16 [int64(int16(c+d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst16 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst16)
			v.AuxInt = int64(int16(c + d))
			return true
		}
		break
	}
	// match: (Add16 <t> (Mul16 x y) (Mul16 x z))
	// result: (Mul16 x (Add16 <t> y z))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMul16 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				y := v_0_1
				if v_1.Op != OpMul16 {
					continue
				}
				_ = v_1.Args[1]
				v_1_0 := v_1.Args[0]
				v_1_1 := v_1.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, v_1_0, v_1_1 = _i2+1, v_1_1, v_1_0 {
					if x != v_1_0 {
						continue
					}
					z := v_1_1
					v.reset(OpMul16)
					v0 := b.NewValue0(v.Pos, OpAdd16, t)
					v0.AddArg2(y, z)
					v.AddArg2(x, v0)
					return true
				}
			}
		}
		break
	}
	// match: (Add16 (Const16 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Add16 (Const16 [1]) (Com16 x))
	// result: (Neg16 x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 || v_0.AuxInt != 1 || v_1.Op != OpCom16 {
				continue
			}
			x := v_1.Args[0]
			v.reset(OpNeg16)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Add16 (Add16 i:(Const16 <t>) z) x)
	// cond: (z.Op != OpConst16 && x.Op != OpConst16)
	// result: (Add16 i (Add16 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAdd16 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst16 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst16 && x.Op != OpConst16) {
					continue
				}
				v.reset(OpAdd16)
				v0 := b.NewValue0(v.Pos, OpAdd16, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Add16 (Sub16 i:(Const16 <t>) z) x)
	// cond: (z.Op != OpConst16 && x.Op != OpConst16)
	// result: (Add16 i (Sub16 <t> x z))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSub16 {
				continue
			}
			z := v_0.Args[1]
			i := v_0.Args[0]
			if i.Op != OpConst16 {
				continue
			}
			t := i.Type
			x := v_1
			if !(z.Op != OpConst16 && x.Op != OpConst16) {
				continue
			}
			v.reset(OpAdd16)
			v0 := b.NewValue0(v.Pos, OpSub16, t)
			v0.AddArg2(x, z)
			v.AddArg2(i, v0)
			return true
		}
		break
	}
	// match: (Add16 (Sub16 z i:(Const16 <t>)) x)
	// cond: (z.Op != OpConst16 && x.Op != OpConst16)
	// result: (Sub16 (Add16 <t> x z) i)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSub16 {
				continue
			}
			_ = v_0.Args[1]
			z := v_0.Args[0]
			i := v_0.Args[1]
			if i.Op != OpConst16 {
				continue
			}
			t := i.Type
			x := v_1
			if !(z.Op != OpConst16 && x.Op != OpConst16) {
				continue
			}
			v.reset(OpSub16)
			v0 := b.NewValue0(v.Pos, OpAdd16, t)
			v0.AddArg2(x, z)
			v.AddArg2(v0, i)
			return true
		}
		break
	}
	// match: (Add16 (Const16 <t> [c]) (Add16 (Const16 <t> [d]) x))
	// result: (Add16 (Const16 <t> [int64(int16(c+d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpAdd16)
				v0 := b.NewValue0(v.Pos, OpConst16, t)
				v0.AuxInt = int64(int16(c + d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Add16 (Const16 <t> [c]) (Sub16 (Const16 <t> [d]) x))
	// result: (Sub16 (Const16 <t> [int64(int16(c+d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpSub16 {
				continue
			}
			x := v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpConst16 || v_1_0.Type != t {
				continue
			}
			d := v_1_0.AuxInt
			v.reset(OpSub16)
			v0 := b.NewValue0(v.Pos, OpConst16, t)
			v0.AuxInt = int64(int16(c + d))
			v.AddArg2(v0, x)
			return true
		}
		break
	}
	// match: (Add16 (Const16 <t> [c]) (Sub16 x (Const16 <t> [d])))
	// result: (Add16 (Const16 <t> [int64(int16(c-d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpSub16 {
				continue
			}
			_ = v_1.Args[1]
			x := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst16 || v_1_1.Type != t {
				continue
			}
			d := v_1_1.AuxInt
			v.reset(OpAdd16)
			v0 := b.NewValue0(v.Pos, OpConst16, t)
			v0.AuxInt = int64(int16(c - d))
			v.AddArg2(v0, x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpAdd32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Add32 (Const32 [c]) (Const32 [d]))
	// result: (Const32 [int64(int32(c+d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst32)
			v.AuxInt = int64(int32(c + d))
			return true
		}
		break
	}
	// match: (Add32 <t> (Mul32 x y) (Mul32 x z))
	// result: (Mul32 x (Add32 <t> y z))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMul32 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				y := v_0_1
				if v_1.Op != OpMul32 {
					continue
				}
				_ = v_1.Args[1]
				v_1_0 := v_1.Args[0]
				v_1_1 := v_1.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, v_1_0, v_1_1 = _i2+1, v_1_1, v_1_0 {
					if x != v_1_0 {
						continue
					}
					z := v_1_1
					v.reset(OpMul32)
					v0 := b.NewValue0(v.Pos, OpAdd32, t)
					v0.AddArg2(y, z)
					v.AddArg2(x, v0)
					return true
				}
			}
		}
		break
	}
	// match: (Add32 (Const32 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Add32 (Const32 [1]) (Com32 x))
	// result: (Neg32 x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != 1 || v_1.Op != OpCom32 {
				continue
			}
			x := v_1.Args[0]
			v.reset(OpNeg32)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Add32 (Add32 i:(Const32 <t>) z) x)
	// cond: (z.Op != OpConst32 && x.Op != OpConst32)
	// result: (Add32 i (Add32 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAdd32 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst32 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst32 && x.Op != OpConst32) {
					continue
				}
				v.reset(OpAdd32)
				v0 := b.NewValue0(v.Pos, OpAdd32, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Add32 (Sub32 i:(Const32 <t>) z) x)
	// cond: (z.Op != OpConst32 && x.Op != OpConst32)
	// result: (Add32 i (Sub32 <t> x z))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSub32 {
				continue
			}
			z := v_0.Args[1]
			i := v_0.Args[0]
			if i.Op != OpConst32 {
				continue
			}
			t := i.Type
			x := v_1
			if !(z.Op != OpConst32 && x.Op != OpConst32) {
				continue
			}
			v.reset(OpAdd32)
			v0 := b.NewValue0(v.Pos, OpSub32, t)
			v0.AddArg2(x, z)
			v.AddArg2(i, v0)
			return true
		}
		break
	}
	// match: (Add32 (Sub32 z i:(Const32 <t>)) x)
	// cond: (z.Op != OpConst32 && x.Op != OpConst32)
	// result: (Sub32 (Add32 <t> x z) i)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSub32 {
				continue
			}
			_ = v_0.Args[1]
			z := v_0.Args[0]
			i := v_0.Args[1]
			if i.Op != OpConst32 {
				continue
			}
			t := i.Type
			x := v_1
			if !(z.Op != OpConst32 && x.Op != OpConst32) {
				continue
			}
			v.reset(OpSub32)
			v0 := b.NewValue0(v.Pos, OpAdd32, t)
			v0.AddArg2(x, z)
			v.AddArg2(v0, i)
			return true
		}
		break
	}
	// match: (Add32 (Const32 <t> [c]) (Add32 (Const32 <t> [d]) x))
	// result: (Add32 (Const32 <t> [int64(int32(c+d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpAdd32)
				v0 := b.NewValue0(v.Pos, OpConst32, t)
				v0.AuxInt = int64(int32(c + d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Add32 (Const32 <t> [c]) (Sub32 (Const32 <t> [d]) x))
	// result: (Sub32 (Const32 <t> [int64(int32(c+d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpSub32 {
				continue
			}
			x := v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpConst32 || v_1_0.Type != t {
				continue
			}
			d := v_1_0.AuxInt
			v.reset(OpSub32)
			v0 := b.NewValue0(v.Pos, OpConst32, t)
			v0.AuxInt = int64(int32(c + d))
			v.AddArg2(v0, x)
			return true
		}
		break
	}
	// match: (Add32 (Const32 <t> [c]) (Sub32 x (Const32 <t> [d])))
	// result: (Add32 (Const32 <t> [int64(int32(c-d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpSub32 {
				continue
			}
			_ = v_1.Args[1]
			x := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst32 || v_1_1.Type != t {
				continue
			}
			d := v_1_1.AuxInt
			v.reset(OpAdd32)
			v0 := b.NewValue0(v.Pos, OpConst32, t)
			v0.AuxInt = int64(int32(c - d))
			v.AddArg2(v0, x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpAdd32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Add32F (Const32F [c]) (Const32F [d]))
	// result: (Const32F [auxFrom32F(auxTo32F(c) + auxTo32F(d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32F {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32F {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst32F)
			v.AuxInt = auxFrom32F(auxTo32F(c) + auxTo32F(d))
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpAdd64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Add64 (Const64 [c]) (Const64 [d]))
	// result: (Const64 [c+d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst64)
			v.AuxInt = c + d
			return true
		}
		break
	}
	// match: (Add64 <t> (Mul64 x y) (Mul64 x z))
	// result: (Mul64 x (Add64 <t> y z))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMul64 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				y := v_0_1
				if v_1.Op != OpMul64 {
					continue
				}
				_ = v_1.Args[1]
				v_1_0 := v_1.Args[0]
				v_1_1 := v_1.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, v_1_0, v_1_1 = _i2+1, v_1_1, v_1_0 {
					if x != v_1_0 {
						continue
					}
					z := v_1_1
					v.reset(OpMul64)
					v0 := b.NewValue0(v.Pos, OpAdd64, t)
					v0.AddArg2(y, z)
					v.AddArg2(x, v0)
					return true
				}
			}
		}
		break
	}
	// match: (Add64 (Const64 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Add64 (Const64 [1]) (Com64 x))
	// result: (Neg64 x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != 1 || v_1.Op != OpCom64 {
				continue
			}
			x := v_1.Args[0]
			v.reset(OpNeg64)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Add64 (Add64 i:(Const64 <t>) z) x)
	// cond: (z.Op != OpConst64 && x.Op != OpConst64)
	// result: (Add64 i (Add64 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAdd64 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst64 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst64 && x.Op != OpConst64) {
					continue
				}
				v.reset(OpAdd64)
				v0 := b.NewValue0(v.Pos, OpAdd64, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Add64 (Sub64 i:(Const64 <t>) z) x)
	// cond: (z.Op != OpConst64 && x.Op != OpConst64)
	// result: (Add64 i (Sub64 <t> x z))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSub64 {
				continue
			}
			z := v_0.Args[1]
			i := v_0.Args[0]
			if i.Op != OpConst64 {
				continue
			}
			t := i.Type
			x := v_1
			if !(z.Op != OpConst64 && x.Op != OpConst64) {
				continue
			}
			v.reset(OpAdd64)
			v0 := b.NewValue0(v.Pos, OpSub64, t)
			v0.AddArg2(x, z)
			v.AddArg2(i, v0)
			return true
		}
		break
	}
	// match: (Add64 (Sub64 z i:(Const64 <t>)) x)
	// cond: (z.Op != OpConst64 && x.Op != OpConst64)
	// result: (Sub64 (Add64 <t> x z) i)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSub64 {
				continue
			}
			_ = v_0.Args[1]
			z := v_0.Args[0]
			i := v_0.Args[1]
			if i.Op != OpConst64 {
				continue
			}
			t := i.Type
			x := v_1
			if !(z.Op != OpConst64 && x.Op != OpConst64) {
				continue
			}
			v.reset(OpSub64)
			v0 := b.NewValue0(v.Pos, OpAdd64, t)
			v0.AddArg2(x, z)
			v.AddArg2(v0, i)
			return true
		}
		break
	}
	// match: (Add64 (Const64 <t> [c]) (Add64 (Const64 <t> [d]) x))
	// result: (Add64 (Const64 <t> [c+d]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpAdd64)
				v0 := b.NewValue0(v.Pos, OpConst64, t)
				v0.AuxInt = c + d
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Add64 (Const64 <t> [c]) (Sub64 (Const64 <t> [d]) x))
	// result: (Sub64 (Const64 <t> [c+d]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpSub64 {
				continue
			}
			x := v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpConst64 || v_1_0.Type != t {
				continue
			}
			d := v_1_0.AuxInt
			v.reset(OpSub64)
			v0 := b.NewValue0(v.Pos, OpConst64, t)
			v0.AuxInt = c + d
			v.AddArg2(v0, x)
			return true
		}
		break
	}
	// match: (Add64 (Const64 <t> [c]) (Sub64 x (Const64 <t> [d])))
	// result: (Add64 (Const64 <t> [c-d]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpSub64 {
				continue
			}
			_ = v_1.Args[1]
			x := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst64 || v_1_1.Type != t {
				continue
			}
			d := v_1_1.AuxInt
			v.reset(OpAdd64)
			v0 := b.NewValue0(v.Pos, OpConst64, t)
			v0.AuxInt = c - d
			v.AddArg2(v0, x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpAdd64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Add64F (Const64F [c]) (Const64F [d]))
	// result: (Const64F [auxFrom64F(auxTo64F(c) + auxTo64F(d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64F {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64F {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst64F)
			v.AuxInt = auxFrom64F(auxTo64F(c) + auxTo64F(d))
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpAdd8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Add8 (Const8 [c]) (Const8 [d]))
	// result: (Const8 [int64(int8(c+d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst8 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst8)
			v.AuxInt = int64(int8(c + d))
			return true
		}
		break
	}
	// match: (Add8 <t> (Mul8 x y) (Mul8 x z))
	// result: (Mul8 x (Add8 <t> y z))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMul8 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				y := v_0_1
				if v_1.Op != OpMul8 {
					continue
				}
				_ = v_1.Args[1]
				v_1_0 := v_1.Args[0]
				v_1_1 := v_1.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, v_1_0, v_1_1 = _i2+1, v_1_1, v_1_0 {
					if x != v_1_0 {
						continue
					}
					z := v_1_1
					v.reset(OpMul8)
					v0 := b.NewValue0(v.Pos, OpAdd8, t)
					v0.AddArg2(y, z)
					v.AddArg2(x, v0)
					return true
				}
			}
		}
		break
	}
	// match: (Add8 (Const8 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Add8 (Const8 [1]) (Com8 x))
	// result: (Neg8 x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 || v_0.AuxInt != 1 || v_1.Op != OpCom8 {
				continue
			}
			x := v_1.Args[0]
			v.reset(OpNeg8)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Add8 (Add8 i:(Const8 <t>) z) x)
	// cond: (z.Op != OpConst8 && x.Op != OpConst8)
	// result: (Add8 i (Add8 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAdd8 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst8 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst8 && x.Op != OpConst8) {
					continue
				}
				v.reset(OpAdd8)
				v0 := b.NewValue0(v.Pos, OpAdd8, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Add8 (Sub8 i:(Const8 <t>) z) x)
	// cond: (z.Op != OpConst8 && x.Op != OpConst8)
	// result: (Add8 i (Sub8 <t> x z))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSub8 {
				continue
			}
			z := v_0.Args[1]
			i := v_0.Args[0]
			if i.Op != OpConst8 {
				continue
			}
			t := i.Type
			x := v_1
			if !(z.Op != OpConst8 && x.Op != OpConst8) {
				continue
			}
			v.reset(OpAdd8)
			v0 := b.NewValue0(v.Pos, OpSub8, t)
			v0.AddArg2(x, z)
			v.AddArg2(i, v0)
			return true
		}
		break
	}
	// match: (Add8 (Sub8 z i:(Const8 <t>)) x)
	// cond: (z.Op != OpConst8 && x.Op != OpConst8)
	// result: (Sub8 (Add8 <t> x z) i)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSub8 {
				continue
			}
			_ = v_0.Args[1]
			z := v_0.Args[0]
			i := v_0.Args[1]
			if i.Op != OpConst8 {
				continue
			}
			t := i.Type
			x := v_1
			if !(z.Op != OpConst8 && x.Op != OpConst8) {
				continue
			}
			v.reset(OpSub8)
			v0 := b.NewValue0(v.Pos, OpAdd8, t)
			v0.AddArg2(x, z)
			v.AddArg2(v0, i)
			return true
		}
		break
	}
	// match: (Add8 (Const8 <t> [c]) (Add8 (Const8 <t> [d]) x))
	// result: (Add8 (Const8 <t> [int64(int8(c+d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst8 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpAdd8)
				v0 := b.NewValue0(v.Pos, OpConst8, t)
				v0.AuxInt = int64(int8(c + d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Add8 (Const8 <t> [c]) (Sub8 (Const8 <t> [d]) x))
	// result: (Sub8 (Const8 <t> [int64(int8(c+d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpSub8 {
				continue
			}
			x := v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpConst8 || v_1_0.Type != t {
				continue
			}
			d := v_1_0.AuxInt
			v.reset(OpSub8)
			v0 := b.NewValue0(v.Pos, OpConst8, t)
			v0.AuxInt = int64(int8(c + d))
			v.AddArg2(v0, x)
			return true
		}
		break
	}
	// match: (Add8 (Const8 <t> [c]) (Sub8 x (Const8 <t> [d])))
	// result: (Add8 (Const8 <t> [int64(int8(c-d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpSub8 {
				continue
			}
			_ = v_1.Args[1]
			x := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst8 || v_1_1.Type != t {
				continue
			}
			d := v_1_1.AuxInt
			v.reset(OpAdd8)
			v0 := b.NewValue0(v.Pos, OpConst8, t)
			v0.AuxInt = int64(int8(c - d))
			v.AddArg2(v0, x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpAddPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AddPtr <t> x (Const64 [c]))
	// result: (OffPtr <t> x [c])
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOffPtr)
		v.Type = t
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (AddPtr <t> x (Const32 [c]))
	// result: (OffPtr <t> x [c])
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOffPtr)
		v.Type = t
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpAnd16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (And16 (Const16 [c]) (Const16 [d]))
	// result: (Const16 [int64(int16(c&d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst16 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst16)
			v.AuxInt = int64(int16(c & d))
			return true
		}
		break
	}
	// match: (And16 (Const16 [m]) (Rsh16Ux64 _ (Const64 [c])))
	// cond: c >= 64-ntz(m)
	// result: (Const16 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			m := v_0.AuxInt
			if v_1.Op != OpRsh16Ux64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst64 {
				continue
			}
			c := v_1_1.AuxInt
			if !(c >= 64-ntz(m)) {
				continue
			}
			v.reset(OpConst16)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And16 (Const16 [m]) (Lsh16x64 _ (Const64 [c])))
	// cond: c >= 64-nlz(m)
	// result: (Const16 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			m := v_0.AuxInt
			if v_1.Op != OpLsh16x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst64 {
				continue
			}
			c := v_1_1.AuxInt
			if !(c >= 64-nlz(m)) {
				continue
			}
			v.reset(OpConst16)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And16 x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (And16 (Const16 [-1]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 || v_0.AuxInt != -1 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (And16 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
				continue
			}
			v.reset(OpConst16)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And16 x (And16 x y))
	// result: (And16 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpAnd16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.reset(OpAnd16)
				v.AddArg2(x, y)
				return true
			}
		}
		break
	}
	// match: (And16 (And16 i:(Const16 <t>) z) x)
	// cond: (z.Op != OpConst16 && x.Op != OpConst16)
	// result: (And16 i (And16 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd16 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst16 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst16 && x.Op != OpConst16) {
					continue
				}
				v.reset(OpAnd16)
				v0 := b.NewValue0(v.Pos, OpAnd16, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (And16 (Const16 <t> [c]) (And16 (Const16 <t> [d]) x))
	// result: (And16 (Const16 <t> [int64(int16(c&d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAnd16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpAnd16)
				v0 := b.NewValue0(v.Pos, OpConst16, t)
				v0.AuxInt = int64(int16(c & d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpAnd32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (And32 (Const32 [c]) (Const32 [d]))
	// result: (Const32 [int64(int32(c&d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst32)
			v.AuxInt = int64(int32(c & d))
			return true
		}
		break
	}
	// match: (And32 (Const32 [m]) (Rsh32Ux64 _ (Const64 [c])))
	// cond: c >= 64-ntz(m)
	// result: (Const32 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			m := v_0.AuxInt
			if v_1.Op != OpRsh32Ux64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst64 {
				continue
			}
			c := v_1_1.AuxInt
			if !(c >= 64-ntz(m)) {
				continue
			}
			v.reset(OpConst32)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And32 (Const32 [m]) (Lsh32x64 _ (Const64 [c])))
	// cond: c >= 64-nlz(m)
	// result: (Const32 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			m := v_0.AuxInt
			if v_1.Op != OpLsh32x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst64 {
				continue
			}
			c := v_1_1.AuxInt
			if !(c >= 64-nlz(m)) {
				continue
			}
			v.reset(OpConst32)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And32 x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (And32 (Const32 [-1]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != -1 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (And32 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
				continue
			}
			v.reset(OpConst32)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And32 x (And32 x y))
	// result: (And32 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpAnd32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.reset(OpAnd32)
				v.AddArg2(x, y)
				return true
			}
		}
		break
	}
	// match: (And32 (And32 i:(Const32 <t>) z) x)
	// cond: (z.Op != OpConst32 && x.Op != OpConst32)
	// result: (And32 i (And32 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd32 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst32 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst32 && x.Op != OpConst32) {
					continue
				}
				v.reset(OpAnd32)
				v0 := b.NewValue0(v.Pos, OpAnd32, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (And32 (Const32 <t> [c]) (And32 (Const32 <t> [d]) x))
	// result: (And32 (Const32 <t> [int64(int32(c&d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAnd32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpAnd32)
				v0 := b.NewValue0(v.Pos, OpConst32, t)
				v0.AuxInt = int64(int32(c & d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpAnd64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (And64 (Const64 [c]) (Const64 [d]))
	// result: (Const64 [c&d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst64)
			v.AuxInt = c & d
			return true
		}
		break
	}
	// match: (And64 (Const64 [m]) (Rsh64Ux64 _ (Const64 [c])))
	// cond: c >= 64-ntz(m)
	// result: (Const64 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			m := v_0.AuxInt
			if v_1.Op != OpRsh64Ux64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst64 {
				continue
			}
			c := v_1_1.AuxInt
			if !(c >= 64-ntz(m)) {
				continue
			}
			v.reset(OpConst64)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And64 (Const64 [m]) (Lsh64x64 _ (Const64 [c])))
	// cond: c >= 64-nlz(m)
	// result: (Const64 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			m := v_0.AuxInt
			if v_1.Op != OpLsh64x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst64 {
				continue
			}
			c := v_1_1.AuxInt
			if !(c >= 64-nlz(m)) {
				continue
			}
			v.reset(OpConst64)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And64 x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (And64 (Const64 [-1]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != -1 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (And64 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
				continue
			}
			v.reset(OpConst64)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And64 x (And64 x y))
	// result: (And64 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpAnd64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.reset(OpAnd64)
				v.AddArg2(x, y)
				return true
			}
		}
		break
	}
	// match: (And64 (And64 i:(Const64 <t>) z) x)
	// cond: (z.Op != OpConst64 && x.Op != OpConst64)
	// result: (And64 i (And64 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd64 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst64 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst64 && x.Op != OpConst64) {
					continue
				}
				v.reset(OpAnd64)
				v0 := b.NewValue0(v.Pos, OpAnd64, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (And64 (Const64 <t> [c]) (And64 (Const64 <t> [d]) x))
	// result: (And64 (Const64 <t> [c&d]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAnd64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpAnd64)
				v0 := b.NewValue0(v.Pos, OpConst64, t)
				v0.AuxInt = c & d
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpAnd8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (And8 (Const8 [c]) (Const8 [d]))
	// result: (Const8 [int64(int8(c&d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst8 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst8)
			v.AuxInt = int64(int8(c & d))
			return true
		}
		break
	}
	// match: (And8 (Const8 [m]) (Rsh8Ux64 _ (Const64 [c])))
	// cond: c >= 64-ntz(m)
	// result: (Const8 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			m := v_0.AuxInt
			if v_1.Op != OpRsh8Ux64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst64 {
				continue
			}
			c := v_1_1.AuxInt
			if !(c >= 64-ntz(m)) {
				continue
			}
			v.reset(OpConst8)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And8 (Const8 [m]) (Lsh8x64 _ (Const64 [c])))
	// cond: c >= 64-nlz(m)
	// result: (Const8 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			m := v_0.AuxInt
			if v_1.Op != OpLsh8x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_1 := v_1.Args[1]
			if v_1_1.Op != OpConst64 {
				continue
			}
			c := v_1_1.AuxInt
			if !(c >= 64-nlz(m)) {
				continue
			}
			v.reset(OpConst8)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And8 x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (And8 (Const8 [-1]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 || v_0.AuxInt != -1 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (And8 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
				continue
			}
			v.reset(OpConst8)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (And8 x (And8 x y))
	// result: (And8 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpAnd8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.reset(OpAnd8)
				v.AddArg2(x, y)
				return true
			}
		}
		break
	}
	// match: (And8 (And8 i:(Const8 <t>) z) x)
	// cond: (z.Op != OpConst8 && x.Op != OpConst8)
	// result: (And8 i (And8 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd8 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst8 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst8 && x.Op != OpConst8) {
					continue
				}
				v.reset(OpAnd8)
				v0 := b.NewValue0(v.Pos, OpAnd8, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (And8 (Const8 <t> [c]) (And8 (Const8 <t> [d]) x))
	// result: (And8 (Const8 <t> [int64(int8(c&d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAnd8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst8 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpAnd8)
				v0 := b.NewValue0(v.Pos, OpConst8, t)
				v0.AuxInt = int64(int8(c & d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpArraySelect(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ArraySelect (ArrayMake1 x))
	// result: x
	for {
		if v_0.Op != OpArrayMake1 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (ArraySelect [0] (IData x))
	// result: (IData x)
	for {
		if v.AuxInt != 0 || v_0.Op != OpIData {
			break
		}
		x := v_0.Args[0]
		v.reset(OpIData)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpCom16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com16 (Com16 x))
	// result: x
	for {
		if v_0.Op != OpCom16 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Com16 (Const16 [c]))
	// result: (Const16 [^c])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = ^c
		return true
	}
	// match: (Com16 (Add16 (Const16 [-1]) x))
	// result: (Neg16 x)
	for {
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst16 || v_0_0.AuxInt != -1 {
				continue
			}
			x := v_0_1
			v.reset(OpNeg16)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpCom32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com32 (Com32 x))
	// result: x
	for {
		if v_0.Op != OpCom32 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Com32 (Const32 [c]))
	// result: (Const32 [^c])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = ^c
		return true
	}
	// match: (Com32 (Add32 (Const32 [-1]) x))
	// result: (Neg32 x)
	for {
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst32 || v_0_0.AuxInt != -1 {
				continue
			}
			x := v_0_1
			v.reset(OpNeg32)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpCom64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com64 (Com64 x))
	// result: x
	for {
		if v_0.Op != OpCom64 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Com64 (Const64 [c]))
	// result: (Const64 [^c])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = ^c
		return true
	}
	// match: (Com64 (Add64 (Const64 [-1]) x))
	// result: (Neg64 x)
	for {
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst64 || v_0_0.AuxInt != -1 {
				continue
			}
			x := v_0_1
			v.reset(OpNeg64)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpCom8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com8 (Com8 x))
	// result: x
	for {
		if v_0.Op != OpCom8 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Com8 (Const8 [c]))
	// result: (Const8 [^c])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst8)
		v.AuxInt = ^c
		return true
	}
	// match: (Com8 (Add8 (Const8 [-1]) x))
	// result: (Neg8 x)
	for {
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst8 || v_0_0.AuxInt != -1 {
				continue
			}
			x := v_0_1
			v.reset(OpNeg8)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpConstInterface(v *Value) bool {
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ConstInterface)
	// result: (IMake (ConstNil <typ.Uintptr>) (ConstNil <typ.BytePtr>))
	for {
		v.reset(OpIMake)
		v0 := b.NewValue0(v.Pos, OpConstNil, typ.Uintptr)
		v1 := b.NewValue0(v.Pos, OpConstNil, typ.BytePtr)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValuegeneric_OpConstSlice(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (ConstSlice)
	// cond: config.PtrSize == 4
	// result: (SliceMake (ConstNil <v.Type.Elem().PtrTo()>) (Const32 <typ.Int> [0]) (Const32 <typ.Int> [0]))
	for {
		if !(config.PtrSize == 4) {
			break
		}
		v.reset(OpSliceMake)
		v0 := b.NewValue0(v.Pos, OpConstNil, v.Type.Elem().PtrTo())
		v1 := b.NewValue0(v.Pos, OpConst32, typ.Int)
		v1.AuxInt = 0
		v2 := b.NewValue0(v.Pos, OpConst32, typ.Int)
		v2.AuxInt = 0
		v.AddArg3(v0, v1, v2)
		return true
	}
	// match: (ConstSlice)
	// cond: config.PtrSize == 8
	// result: (SliceMake (ConstNil <v.Type.Elem().PtrTo()>) (Const64 <typ.Int> [0]) (Const64 <typ.Int> [0]))
	for {
		if !(config.PtrSize == 8) {
			break
		}
		v.reset(OpSliceMake)
		v0 := b.NewValue0(v.Pos, OpConstNil, v.Type.Elem().PtrTo())
		v1 := b.NewValue0(v.Pos, OpConst64, typ.Int)
		v1.AuxInt = 0
		v2 := b.NewValue0(v.Pos, OpConst64, typ.Int)
		v2.AuxInt = 0
		v.AddArg3(v0, v1, v2)
		return true
	}
	return false
}
func rewriteValuegeneric_OpConstString(v *Value) bool {
	b := v.Block
	config := b.Func.Config
	fe := b.Func.fe
	typ := &b.Func.Config.Types
	// match: (ConstString {s})
	// cond: config.PtrSize == 4 && s.(string) == ""
	// result: (StringMake (ConstNil) (Const32 <typ.Int> [0]))
	for {
		s := v.Aux
		if !(config.PtrSize == 4 && s.(string) == "") {
			break
		}
		v.reset(OpStringMake)
		v0 := b.NewValue0(v.Pos, OpConstNil, typ.BytePtr)
		v1 := b.NewValue0(v.Pos, OpConst32, typ.Int)
		v1.AuxInt = 0
		v.AddArg2(v0, v1)
		return true
	}
	// match: (ConstString {s})
	// cond: config.PtrSize == 8 && s.(string) == ""
	// result: (StringMake (ConstNil) (Const64 <typ.Int> [0]))
	for {
		s := v.Aux
		if !(config.PtrSize == 8 && s.(string) == "") {
			break
		}
		v.reset(OpStringMake)
		v0 := b.NewValue0(v.Pos, OpConstNil, typ.BytePtr)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.Int)
		v1.AuxInt = 0
		v.AddArg2(v0, v1)
		return true
	}
	// match: (ConstString {s})
	// cond: config.PtrSize == 4 && s.(string) != ""
	// result: (StringMake (Addr <typ.BytePtr> {fe.StringData(s.(string))} (SB)) (Const32 <typ.Int> [int64(len(s.(string)))]))
	for {
		s := v.Aux
		if !(config.PtrSize == 4 && s.(string) != "") {
			break
		}
		v.reset(OpStringMake)
		v0 := b.NewValue0(v.Pos, OpAddr, typ.BytePtr)
		v0.Aux = fe.StringData(s.(string))
		v1 := b.NewValue0(v.Pos, OpSB, typ.Uintptr)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.Int)
		v2.AuxInt = int64(len(s.(string)))
		v.AddArg2(v0, v2)
		return true
	}
	// match: (ConstString {s})
	// cond: config.PtrSize == 8 && s.(string) != ""
	// result: (StringMake (Addr <typ.BytePtr> {fe.StringData(s.(string))} (SB)) (Const64 <typ.Int> [int64(len(s.(string)))]))
	for {
		s := v.Aux
		if !(config.PtrSize == 8 && s.(string) != "") {
			break
		}
		v.reset(OpStringMake)
		v0 := b.NewValue0(v.Pos, OpAddr, typ.BytePtr)
		v0.Aux = fe.StringData(s.(string))
		v1 := b.NewValue0(v.Pos, OpSB, typ.Uintptr)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.Int)
		v2.AuxInt = int64(len(s.(string)))
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValuegeneric_OpConvert(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Convert (Add64 (Convert ptr mem) off) mem)
	// result: (Add64 ptr off)
	for {
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConvert {
				continue
			}
			mem := v_0_0.Args[1]
			ptr := v_0_0.Args[0]
			off := v_0_1
			if mem != v_1 {
				continue
			}
			v.reset(OpAdd64)
			v.AddArg2(ptr, off)
			return true
		}
		break
	}
	// match: (Convert (Add32 (Convert ptr mem) off) mem)
	// result: (Add32 ptr off)
	for {
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConvert {
				continue
			}
			mem := v_0_0.Args[1]
			ptr := v_0_0.Args[0]
			off := v_0_1
			if mem != v_1 {
				continue
			}
			v.reset(OpAdd32)
			v.AddArg2(ptr, off)
			return true
		}
		break
	}
	// match: (Convert (Convert ptr mem) mem)
	// result: ptr
	for {
		if v_0.Op != OpConvert {
			break
		}
		mem := v_0.Args[1]
		ptr := v_0.Args[0]
		if mem != v_1 {
			break
		}
		v.copyOf(ptr)
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt32Fto32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32Fto32 (Const32F [c]))
	// result: (Const32 [int64(int32(auxTo32F(c)))])
	for {
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(auxTo32F(c)))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt32Fto64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32Fto64 (Const32F [c]))
	// result: (Const64 [int64(auxTo32F(c))])
	for {
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(auxTo32F(c))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt32Fto64F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32Fto64F (Const32F [c]))
	// result: (Const64F [c])
	for {
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt32to32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32to32F (Const32 [c]))
	// result: (Const32F [auxFrom32F(float32(int32(c)))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = auxFrom32F(float32(int32(c)))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt32to64F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32to64F (Const32 [c]))
	// result: (Const64F [auxFrom64F(float64(int32(c)))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = auxFrom64F(float64(int32(c)))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt64Fto32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt64Fto32 (Const64F [c]))
	// result: (Const32 [int64(int32(auxTo64F(c)))])
	for {
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(auxTo64F(c)))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt64Fto32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt64Fto32F (Const64F [c]))
	// result: (Const32F [auxFrom32F(float32(auxTo64F(c)))])
	for {
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = auxFrom32F(float32(auxTo64F(c)))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt64Fto64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt64Fto64 (Const64F [c]))
	// result: (Const64 [int64(auxTo64F(c))])
	for {
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(auxTo64F(c))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt64to32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt64to32F (Const64 [c]))
	// result: (Const32F [auxFrom32F(float32(c))])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = auxFrom32F(float32(c))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvt64to64F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt64to64F (Const64 [c]))
	// result: (Const64F [auxFrom64F(float64(c))])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = auxFrom64F(float64(c))
		return true
	}
	return false
}
func rewriteValuegeneric_OpCvtBoolToUint8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CvtBoolToUint8 (ConstBool [c]))
	// result: (Const8 [c])
	for {
		if v_0.Op != OpConstBool {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst8)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 (Const16 [c]) (Const16 [d]))
	// cond: d != 0
	// result: (Const16 [int64(int16(c)/int16(d))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c) / int16(d))
		return true
	}
	// match: (Div16 n (Const16 [c]))
	// cond: isNonNegative(n) && isPowerOfTwo(c&0xffff)
	// result: (Rsh16Ux64 n (Const64 <typ.UInt64> [log2(c&0xffff)]))
	for {
		n := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xffff)) {
			break
		}
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xffff)
		v.AddArg2(n, v0)
		return true
	}
	// match: (Div16 <t> n (Const16 [c]))
	// cond: c < 0 && c != -1<<15
	// result: (Neg16 (Div16 <t> n (Const16 <t> [-c])))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<15) {
			break
		}
		v.reset(OpNeg16)
		v0 := b.NewValue0(v.Pos, OpDiv16, t)
		v1 := b.NewValue0(v.Pos, OpConst16, t)
		v1.AuxInt = -c
		v0.AddArg2(n, v1)
		v.AddArg(v0)
		return true
	}
	// match: (Div16 <t> x (Const16 [-1<<15]))
	// result: (Rsh16Ux64 (And16 <t> x (Neg16 <t> x)) (Const64 <typ.UInt64> [15]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 || v_1.AuxInt != -1<<15 {
			break
		}
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpAnd16, t)
		v1 := b.NewValue0(v.Pos, OpNeg16, t)
		v1.AddArg(x)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = 15
		v.AddArg2(v0, v2)
		return true
	}
	// match: (Div16 <t> n (Const16 [c]))
	// cond: isPowerOfTwo(c)
	// result: (Rsh16x64 (Add16 <t> n (Rsh16Ux64 <t> (Rsh16x64 <t> n (Const64 <typ.UInt64> [15])) (Const64 <typ.UInt64> [16-log2(c)]))) (Const64 <typ.UInt64> [log2(c)]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh16x64)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v1 := b.NewValue0(v.Pos, OpRsh16Ux64, t)
		v2 := b.NewValue0(v.Pos, OpRsh16x64, t)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 15
		v2.AddArg2(n, v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 16 - log2(c)
		v1.AddArg2(v2, v4)
		v0.AddArg2(n, v1)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = log2(c)
		v.AddArg2(v0, v5)
		return true
	}
	// match: (Div16 <t> x (Const16 [c]))
	// cond: smagicOK(16,c)
	// result: (Sub16 <t> (Rsh32x64 <t> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(smagic(16,c).m)]) (SignExt16to32 x)) (Const64 <typ.UInt64> [16+smagic(16,c).s])) (Rsh32x64 <t> (SignExt16to32 x) (Const64 <typ.UInt64> [31])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(16, c)) {
			break
		}
		v.reset(OpSub16)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(smagic(16, c).m)
		v3 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v3.AddArg(x)
		v1.AddArg2(v2, v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 16 + smagic(16, c).s
		v0.AddArg2(v1, v4)
		v5 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v6 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v6.AddArg(x)
		v7 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v7.AuxInt = 31
		v5.AddArg2(v6, v7)
		v.AddArg2(v0, v5)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Div16u (Const16 [c]) (Const16 [d]))
	// cond: d != 0
	// result: (Const16 [int64(int16(uint16(c)/uint16(d)))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = int64(int16(uint16(c) / uint16(d)))
		return true
	}
	// match: (Div16u n (Const16 [c]))
	// cond: isPowerOfTwo(c&0xffff)
	// result: (Rsh16Ux64 n (Const64 <typ.UInt64> [log2(c&0xffff)]))
	for {
		n := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xffff)) {
			break
		}
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xffff)
		v.AddArg2(n, v0)
		return true
	}
	// match: (Div16u x (Const16 [c]))
	// cond: umagicOK(16, c) && config.RegSize == 8
	// result: (Trunc64to16 (Rsh64Ux64 <typ.UInt64> (Mul64 <typ.UInt64> (Const64 <typ.UInt64> [int64(1<<16+umagic(16,c).m)]) (ZeroExt16to64 x)) (Const64 <typ.UInt64> [16+umagic(16,c).s])))
	for {
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(16, c) && config.RegSize == 8) {
			break
		}
		v.reset(OpTrunc64to16)
		v0 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(1<<16 + umagic(16, c).m)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(x)
		v1.AddArg2(v2, v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 16 + umagic(16, c).s
		v0.AddArg2(v1, v4)
		v.AddArg(v0)
		return true
	}
	// match: (Div16u x (Const16 [c]))
	// cond: umagicOK(16, c) && config.RegSize == 4 && umagic(16,c).m&1 == 0
	// result: (Trunc32to16 (Rsh32Ux64 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(1<<15+umagic(16,c).m/2)]) (ZeroExt16to32 x)) (Const64 <typ.UInt64> [16+umagic(16,c).s-1])))
	for {
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(16, c) && config.RegSize == 4 && umagic(16, c).m&1 == 0) {
			break
		}
		v.reset(OpTrunc32to16)
		v0 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(1<<15 + umagic(16, c).m/2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(x)
		v1.AddArg2(v2, v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 16 + umagic(16, c).s - 1
		v0.AddArg2(v1, v4)
		v.AddArg(v0)
		return true
	}
	// match: (Div16u x (Const16 [c]))
	// cond: umagicOK(16, c) && config.RegSize == 4 && c&1 == 0
	// result: (Trunc32to16 (Rsh32Ux64 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(1<<15+(umagic(16,c).m+1)/2)]) (Rsh32Ux64 <typ.UInt32> (ZeroExt16to32 x) (Const64 <typ.UInt64> [1]))) (Const64 <typ.UInt64> [16+umagic(16,c).s-2])))
	for {
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(16, c) && config.RegSize == 4 && c&1 == 0) {
			break
		}
		v.reset(OpTrunc32to16)
		v0 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(1<<15 + (umagic(16, c).m+1)/2)
		v3 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = 1
		v3.AddArg2(v4, v5)
		v1.AddArg2(v2, v3)
		v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v6.AuxInt = 16 + umagic(16, c).s - 2
		v0.AddArg2(v1, v6)
		v.AddArg(v0)
		return true
	}
	// match: (Div16u x (Const16 [c]))
	// cond: umagicOK(16, c) && config.RegSize == 4 && config.useAvg
	// result: (Trunc32to16 (Rsh32Ux64 <typ.UInt32> (Avg32u (Lsh32x64 <typ.UInt32> (ZeroExt16to32 x) (Const64 <typ.UInt64> [16])) (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(umagic(16,c).m)]) (ZeroExt16to32 x))) (Const64 <typ.UInt64> [16+umagic(16,c).s-1])))
	for {
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(16, c) && config.RegSize == 4 && config.useAvg) {
			break
		}
		v.reset(OpTrunc32to16)
		v0 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpAvg32u, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpLsh32x64, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(x)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 16
		v2.AddArg2(v3, v4)
		v5 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v6 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v6.AuxInt = int64(umagic(16, c).m)
		v7 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v7.AddArg(x)
		v5.AddArg2(v6, v7)
		v1.AddArg2(v2, v5)
		v8 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v8.AuxInt = 16 + umagic(16, c).s - 1
		v0.AddArg2(v1, v8)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Div32 (Const32 [c]) (Const32 [d]))
	// cond: d != 0
	// result: (Const32 [int64(int32(c)/int32(d))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c) / int32(d))
		return true
	}
	// match: (Div32 n (Const32 [c]))
	// cond: isNonNegative(n) && isPowerOfTwo(c&0xffffffff)
	// result: (Rsh32Ux64 n (Const64 <typ.UInt64> [log2(c&0xffffffff)]))
	for {
		n := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xffffffff)) {
			break
		}
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xffffffff)
		v.AddArg2(n, v0)
		return true
	}
	// match: (Div32 <t> n (Const32 [c]))
	// cond: c < 0 && c != -1<<31
	// result: (Neg32 (Div32 <t> n (Const32 <t> [-c])))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<31) {
			break
		}
		v.reset(OpNeg32)
		v0 := b.NewValue0(v.Pos, OpDiv32, t)
		v1 := b.NewValue0(v.Pos, OpConst32, t)
		v1.AuxInt = -c
		v0.AddArg2(n, v1)
		v.AddArg(v0)
		return true
	}
	// match: (Div32 <t> x (Const32 [-1<<31]))
	// result: (Rsh32Ux64 (And32 <t> x (Neg32 <t> x)) (Const64 <typ.UInt64> [31]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 || v_1.AuxInt != -1<<31 {
			break
		}
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpAnd32, t)
		v1 := b.NewValue0(v.Pos, OpNeg32, t)
		v1.AddArg(x)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = 31
		v.AddArg2(v0, v2)
		return true
	}
	// match: (Div32 <t> n (Const32 [c]))
	// cond: isPowerOfTwo(c)
	// result: (Rsh32x64 (Add32 <t> n (Rsh32Ux64 <t> (Rsh32x64 <t> n (Const64 <typ.UInt64> [31])) (Const64 <typ.UInt64> [32-log2(c)]))) (Const64 <typ.UInt64> [log2(c)]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh32x64)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v1 := b.NewValue0(v.Pos, OpRsh32Ux64, t)
		v2 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 31
		v2.AddArg2(n, v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 32 - log2(c)
		v1.AddArg2(v2, v4)
		v0.AddArg2(n, v1)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = log2(c)
		v.AddArg2(v0, v5)
		return true
	}
	// match: (Div32 <t> x (Const32 [c]))
	// cond: smagicOK(32,c) && config.RegSize == 8
	// result: (Sub32 <t> (Rsh64x64 <t> (Mul64 <typ.UInt64> (Const64 <typ.UInt64> [int64(smagic(32,c).m)]) (SignExt32to64 x)) (Const64 <typ.UInt64> [32+smagic(32,c).s])) (Rsh64x64 <t> (SignExt32to64 x) (Const64 <typ.UInt64> [63])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(32, c) && config.RegSize == 8) {
			break
		}
		v.reset(OpSub32)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v1 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(smagic(32, c).m)
		v3 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v3.AddArg(x)
		v1.AddArg2(v2, v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 32 + smagic(32, c).s
		v0.AddArg2(v1, v4)
		v5 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v6 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v6.AddArg(x)
		v7 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v7.AuxInt = 63
		v5.AddArg2(v6, v7)
		v.AddArg2(v0, v5)
		return true
	}
	// match: (Div32 <t> x (Const32 [c]))
	// cond: smagicOK(32,c) && config.RegSize == 4 && smagic(32,c).m&1 == 0 && config.useHmul
	// result: (Sub32 <t> (Rsh32x64 <t> (Hmul32 <t> (Const32 <typ.UInt32> [int64(int32(smagic(32,c).m/2))]) x) (Const64 <typ.UInt64> [smagic(32,c).s-1])) (Rsh32x64 <t> x (Const64 <typ.UInt64> [31])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(32, c) && config.RegSize == 4 && smagic(32, c).m&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpSub32)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v1 := b.NewValue0(v.Pos, OpHmul32, t)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(int32(smagic(32, c).m / 2))
		v1.AddArg2(v2, x)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = smagic(32, c).s - 1
		v0.AddArg2(v1, v3)
		v4 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = 31
		v4.AddArg2(x, v5)
		v.AddArg2(v0, v4)
		return true
	}
	// match: (Div32 <t> x (Const32 [c]))
	// cond: smagicOK(32,c) && config.RegSize == 4 && smagic(32,c).m&1 != 0 && config.useHmul
	// result: (Sub32 <t> (Rsh32x64 <t> (Add32 <t> (Hmul32 <t> (Const32 <typ.UInt32> [int64(int32(smagic(32,c).m))]) x) x) (Const64 <typ.UInt64> [smagic(32,c).s])) (Rsh32x64 <t> x (Const64 <typ.UInt64> [31])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(32, c) && config.RegSize == 4 && smagic(32, c).m&1 != 0 && config.useHmul) {
			break
		}
		v.reset(OpSub32)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v1 := b.NewValue0(v.Pos, OpAdd32, t)
		v2 := b.NewValue0(v.Pos, OpHmul32, t)
		v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v3.AuxInt = int64(int32(smagic(32, c).m))
		v2.AddArg2(v3, x)
		v1.AddArg2(v2, x)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = smagic(32, c).s
		v0.AddArg2(v1, v4)
		v5 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v6.AuxInt = 31
		v5.AddArg2(x, v6)
		v.AddArg2(v0, v5)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Div32F (Const32F [c]) (Const32F [d]))
	// result: (Const32F [auxFrom32F(auxTo32F(c) / auxTo32F(d))])
	for {
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = auxFrom32F(auxTo32F(c) / auxTo32F(d))
		return true
	}
	// match: (Div32F x (Const32F <t> [c]))
	// cond: reciprocalExact32(auxTo32F(c))
	// result: (Mul32F x (Const32F <t> [auxFrom32F(1/auxTo32F(c))]))
	for {
		x := v_0
		if v_1.Op != OpConst32F {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(reciprocalExact32(auxTo32F(c))) {
			break
		}
		v.reset(OpMul32F)
		v0 := b.NewValue0(v.Pos, OpConst32F, t)
		v0.AuxInt = auxFrom32F(1 / auxTo32F(c))
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Div32u (Const32 [c]) (Const32 [d]))
	// cond: d != 0
	// result: (Const32 [int64(int32(uint32(c)/uint32(d)))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int64(int32(uint32(c) / uint32(d)))
		return true
	}
	// match: (Div32u n (Const32 [c]))
	// cond: isPowerOfTwo(c&0xffffffff)
	// result: (Rsh32Ux64 n (Const64 <typ.UInt64> [log2(c&0xffffffff)]))
	for {
		n := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xffffffff)) {
			break
		}
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xffffffff)
		v.AddArg2(n, v0)
		return true
	}
	// match: (Div32u x (Const32 [c]))
	// cond: umagicOK(32, c) && config.RegSize == 4 && umagic(32,c).m&1 == 0 && config.useHmul
	// result: (Rsh32Ux64 <typ.UInt32> (Hmul32u <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(1<<31+umagic(32,c).m/2))]) x) (Const64 <typ.UInt64> [umagic(32,c).s-1]))
	for {
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 4 && umagic(32, c).m&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpRsh32Ux64)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpHmul32u, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v1.AuxInt = int64(int32(1<<31 + umagic(32, c).m/2))
		v0.AddArg2(v1, x)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = umagic(32, c).s - 1
		v.AddArg2(v0, v2)
		return true
	}
	// match: (Div32u x (Const32 [c]))
	// cond: umagicOK(32, c) && config.RegSize == 4 && c&1 == 0 && config.useHmul
	// result: (Rsh32Ux64 <typ.UInt32> (Hmul32u <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(1<<31+(umagic(32,c).m+1)/2))]) (Rsh32Ux64 <typ.UInt32> x (Const64 <typ.UInt64> [1]))) (Const64 <typ.UInt64> [umagic(32,c).s-2]))
	for {
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 4 && c&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpRsh32Ux64)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpHmul32u, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v1.AuxInt = int64(int32(1<<31 + (umagic(32, c).m+1)/2))
		v2 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 1
		v2.AddArg2(x, v3)
		v0.AddArg2(v1, v2)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = umagic(32, c).s - 2
		v.AddArg2(v0, v4)
		return true
	}
	// match: (Div32u x (Const32 [c]))
	// cond: umagicOK(32, c) && config.RegSize == 4 && config.useAvg && config.useHmul
	// result: (Rsh32Ux64 <typ.UInt32> (Avg32u x (Hmul32u <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(umagic(32,c).m))]) x)) (Const64 <typ.UInt64> [umagic(32,c).s-1]))
	for {
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 4 && config.useAvg && config.useHmul) {
			break
		}
		v.reset(OpRsh32Ux64)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpAvg32u, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpHmul32u, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(int32(umagic(32, c).m))
		v1.AddArg2(v2, x)
		v0.AddArg2(x, v1)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = umagic(32, c).s - 1
		v.AddArg2(v0, v3)
		return true
	}
	// match: (Div32u x (Const32 [c]))
	// cond: umagicOK(32, c) && config.RegSize == 8 && umagic(32,c).m&1 == 0
	// result: (Trunc64to32 (Rsh64Ux64 <typ.UInt64> (Mul64 <typ.UInt64> (Const64 <typ.UInt64> [int64(1<<31+umagic(32,c).m/2)]) (ZeroExt32to64 x)) (Const64 <typ.UInt64> [32+umagic(32,c).s-1])))
	for {
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 8 && umagic(32, c).m&1 == 0) {
			break
		}
		v.reset(OpTrunc64to32)
		v0 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(1<<31 + umagic(32, c).m/2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(x)
		v1.AddArg2(v2, v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 32 + umagic(32, c).s - 1
		v0.AddArg2(v1, v4)
		v.AddArg(v0)
		return true
	}
	// match: (Div32u x (Const32 [c]))
	// cond: umagicOK(32, c) && config.RegSize == 8 && c&1 == 0
	// result: (Trunc64to32 (Rsh64Ux64 <typ.UInt64> (Mul64 <typ.UInt64> (Const64 <typ.UInt64> [int64(1<<31+(umagic(32,c).m+1)/2)]) (Rsh64Ux64 <typ.UInt64> (ZeroExt32to64 x) (Const64 <typ.UInt64> [1]))) (Const64 <typ.UInt64> [32+umagic(32,c).s-2])))
	for {
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 8 && c&1 == 0) {
			break
		}
		v.reset(OpTrunc64to32)
		v0 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(1<<31 + (umagic(32, c).m+1)/2)
		v3 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = 1
		v3.AddArg2(v4, v5)
		v1.AddArg2(v2, v3)
		v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v6.AuxInt = 32 + umagic(32, c).s - 2
		v0.AddArg2(v1, v6)
		v.AddArg(v0)
		return true
	}
	// match: (Div32u x (Const32 [c]))
	// cond: umagicOK(32, c) && config.RegSize == 8 && config.useAvg
	// result: (Trunc64to32 (Rsh64Ux64 <typ.UInt64> (Avg64u (Lsh64x64 <typ.UInt64> (ZeroExt32to64 x) (Const64 <typ.UInt64> [32])) (Mul64 <typ.UInt64> (Const64 <typ.UInt32> [int64(umagic(32,c).m)]) (ZeroExt32to64 x))) (Const64 <typ.UInt64> [32+umagic(32,c).s-1])))
	for {
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(32, c) && config.RegSize == 8 && config.useAvg) {
			break
		}
		v.reset(OpTrunc64to32)
		v0 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpAvg64u, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpLsh64x64, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(x)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 32
		v2.AddArg2(v3, v4)
		v5 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
		v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt32)
		v6.AuxInt = int64(umagic(32, c).m)
		v7 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v7.AddArg(x)
		v5.AddArg2(v6, v7)
		v1.AddArg2(v2, v5)
		v8 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v8.AuxInt = 32 + umagic(32, c).s - 1
		v0.AddArg2(v1, v8)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Div64 (Const64 [c]) (Const64 [d]))
	// cond: d != 0
	// result: (Const64 [c/d])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = c / d
		return true
	}
	// match: (Div64 n (Const64 [c]))
	// cond: isNonNegative(n) && isPowerOfTwo(c)
	// result: (Rsh64Ux64 n (Const64 <typ.UInt64> [log2(c)]))
	for {
		n := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg2(n, v0)
		return true
	}
	// match: (Div64 n (Const64 [-1<<63]))
	// cond: isNonNegative(n)
	// result: (Const64 [0])
	for {
		n := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != -1<<63 || !(isNonNegative(n)) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	// match: (Div64 <t> n (Const64 [c]))
	// cond: c < 0 && c != -1<<63
	// result: (Neg64 (Div64 <t> n (Const64 <t> [-c])))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<63) {
			break
		}
		v.reset(OpNeg64)
		v0 := b.NewValue0(v.Pos, OpDiv64, t)
		v1 := b.NewValue0(v.Pos, OpConst64, t)
		v1.AuxInt = -c
		v0.AddArg2(n, v1)
		v.AddArg(v0)
		return true
	}
	// match: (Div64 <t> x (Const64 [-1<<63]))
	// result: (Rsh64Ux64 (And64 <t> x (Neg64 <t> x)) (Const64 <typ.UInt64> [63]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != -1<<63 {
			break
		}
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpAnd64, t)
		v1 := b.NewValue0(v.Pos, OpNeg64, t)
		v1.AddArg(x)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = 63
		v.AddArg2(v0, v2)
		return true
	}
	// match: (Div64 <t> n (Const64 [c]))
	// cond: isPowerOfTwo(c)
	// result: (Rsh64x64 (Add64 <t> n (Rsh64Ux64 <t> (Rsh64x64 <t> n (Const64 <typ.UInt64> [63])) (Const64 <typ.UInt64> [64-log2(c)]))) (Const64 <typ.UInt64> [log2(c)]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v1 := b.NewValue0(v.Pos, OpRsh64Ux64, t)
		v2 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 63
		v2.AddArg2(n, v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 64 - log2(c)
		v1.AddArg2(v2, v4)
		v0.AddArg2(n, v1)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = log2(c)
		v.AddArg2(v0, v5)
		return true
	}
	// match: (Div64 <t> x (Const64 [c]))
	// cond: smagicOK(64,c) && smagic(64,c).m&1 == 0 && config.useHmul
	// result: (Sub64 <t> (Rsh64x64 <t> (Hmul64 <t> (Const64 <typ.UInt64> [int64(smagic(64,c).m/2)]) x) (Const64 <typ.UInt64> [smagic(64,c).s-1])) (Rsh64x64 <t> x (Const64 <typ.UInt64> [63])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(64, c) && smagic(64, c).m&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpSub64)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v1 := b.NewValue0(v.Pos, OpHmul64, t)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(smagic(64, c).m / 2)
		v1.AddArg2(v2, x)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = smagic(64, c).s - 1
		v0.AddArg2(v1, v3)
		v4 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = 63
		v4.AddArg2(x, v5)
		v.AddArg2(v0, v4)
		return true
	}
	// match: (Div64 <t> x (Const64 [c]))
	// cond: smagicOK(64,c) && smagic(64,c).m&1 != 0 && config.useHmul
	// result: (Sub64 <t> (Rsh64x64 <t> (Add64 <t> (Hmul64 <t> (Const64 <typ.UInt64> [int64(smagic(64,c).m)]) x) x) (Const64 <typ.UInt64> [smagic(64,c).s])) (Rsh64x64 <t> x (Const64 <typ.UInt64> [63])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(64, c) && smagic(64, c).m&1 != 0 && config.useHmul) {
			break
		}
		v.reset(OpSub64)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v1 := b.NewValue0(v.Pos, OpAdd64, t)
		v2 := b.NewValue0(v.Pos, OpHmul64, t)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = int64(smagic(64, c).m)
		v2.AddArg2(v3, x)
		v1.AddArg2(v2, x)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = smagic(64, c).s
		v0.AddArg2(v1, v4)
		v5 := b.NewValue0(v.Pos, OpRsh64x64, t)
		v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v6.AuxInt = 63
		v5.AddArg2(x, v6)
		v.AddArg2(v0, v5)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Div64F (Const64F [c]) (Const64F [d]))
	// result: (Const64F [auxFrom64F(auxTo64F(c) / auxTo64F(d))])
	for {
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = auxFrom64F(auxTo64F(c) / auxTo64F(d))
		return true
	}
	// match: (Div64F x (Const64F <t> [c]))
	// cond: reciprocalExact64(auxTo64F(c))
	// result: (Mul64F x (Const64F <t> [auxFrom64F(1/auxTo64F(c))]))
	for {
		x := v_0
		if v_1.Op != OpConst64F {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(reciprocalExact64(auxTo64F(c))) {
			break
		}
		v.reset(OpMul64F)
		v0 := b.NewValue0(v.Pos, OpConst64F, t)
		v0.AuxInt = auxFrom64F(1 / auxTo64F(c))
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Div64u (Const64 [c]) (Const64 [d]))
	// cond: d != 0
	// result: (Const64 [int64(uint64(c)/uint64(d))])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = int64(uint64(c) / uint64(d))
		return true
	}
	// match: (Div64u n (Const64 [c]))
	// cond: isPowerOfTwo(c)
	// result: (Rsh64Ux64 n (Const64 <typ.UInt64> [log2(c)]))
	for {
		n := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c)
		v.AddArg2(n, v0)
		return true
	}
	// match: (Div64u n (Const64 [-1<<63]))
	// result: (Rsh64Ux64 n (Const64 <typ.UInt64> [63]))
	for {
		n := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != -1<<63 {
			break
		}
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = 63
		v.AddArg2(n, v0)
		return true
	}
	// match: (Div64u x (Const64 [c]))
	// cond: umagicOK(64, c) && config.RegSize == 8 && umagic(64,c).m&1 == 0 && config.useHmul
	// result: (Rsh64Ux64 <typ.UInt64> (Hmul64u <typ.UInt64> (Const64 <typ.UInt64> [int64(1<<63+umagic(64,c).m/2)]) x) (Const64 <typ.UInt64> [umagic(64,c).s-1]))
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(64, c) && config.RegSize == 8 && umagic(64, c).m&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpRsh64Ux64)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpHmul64u, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = int64(1<<63 + umagic(64, c).m/2)
		v0.AddArg2(v1, x)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = umagic(64, c).s - 1
		v.AddArg2(v0, v2)
		return true
	}
	// match: (Div64u x (Const64 [c]))
	// cond: umagicOK(64, c) && config.RegSize == 8 && c&1 == 0 && config.useHmul
	// result: (Rsh64Ux64 <typ.UInt64> (Hmul64u <typ.UInt64> (Const64 <typ.UInt64> [int64(1<<63+(umagic(64,c).m+1)/2)]) (Rsh64Ux64 <typ.UInt64> x (Const64 <typ.UInt64> [1]))) (Const64 <typ.UInt64> [umagic(64,c).s-2]))
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(64, c) && config.RegSize == 8 && c&1 == 0 && config.useHmul) {
			break
		}
		v.reset(OpRsh64Ux64)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpHmul64u, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v1.AuxInt = int64(1<<63 + (umagic(64, c).m+1)/2)
		v2 := b.NewValue0(v.Pos, OpRsh64Ux64, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 1
		v2.AddArg2(x, v3)
		v0.AddArg2(v1, v2)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = umagic(64, c).s - 2
		v.AddArg2(v0, v4)
		return true
	}
	// match: (Div64u x (Const64 [c]))
	// cond: umagicOK(64, c) && config.RegSize == 8 && config.useAvg && config.useHmul
	// result: (Rsh64Ux64 <typ.UInt64> (Avg64u x (Hmul64u <typ.UInt64> (Const64 <typ.UInt64> [int64(umagic(64,c).m)]) x)) (Const64 <typ.UInt64> [umagic(64,c).s-1]))
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(64, c) && config.RegSize == 8 && config.useAvg && config.useHmul) {
			break
		}
		v.reset(OpRsh64Ux64)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpAvg64u, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpHmul64u, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = int64(umagic(64, c).m)
		v1.AddArg2(v2, x)
		v0.AddArg2(x, v1)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = umagic(64, c).s - 1
		v.AddArg2(v0, v3)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 (Const8 [c]) (Const8 [d]))
	// cond: d != 0
	// result: (Const8 [int64(int8(c)/int8(d))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c) / int8(d))
		return true
	}
	// match: (Div8 n (Const8 [c]))
	// cond: isNonNegative(n) && isPowerOfTwo(c&0xff)
	// result: (Rsh8Ux64 n (Const64 <typ.UInt64> [log2(c&0xff)]))
	for {
		n := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xff)) {
			break
		}
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xff)
		v.AddArg2(n, v0)
		return true
	}
	// match: (Div8 <t> n (Const8 [c]))
	// cond: c < 0 && c != -1<<7
	// result: (Neg8 (Div8 <t> n (Const8 <t> [-c])))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<7) {
			break
		}
		v.reset(OpNeg8)
		v0 := b.NewValue0(v.Pos, OpDiv8, t)
		v1 := b.NewValue0(v.Pos, OpConst8, t)
		v1.AuxInt = -c
		v0.AddArg2(n, v1)
		v.AddArg(v0)
		return true
	}
	// match: (Div8 <t> x (Const8 [-1<<7 ]))
	// result: (Rsh8Ux64 (And8 <t> x (Neg8 <t> x)) (Const64 <typ.UInt64> [7 ]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 || v_1.AuxInt != -1<<7 {
			break
		}
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpAnd8, t)
		v1 := b.NewValue0(v.Pos, OpNeg8, t)
		v1.AddArg(x)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v2.AuxInt = 7
		v.AddArg2(v0, v2)
		return true
	}
	// match: (Div8 <t> n (Const8 [c]))
	// cond: isPowerOfTwo(c)
	// result: (Rsh8x64 (Add8 <t> n (Rsh8Ux64 <t> (Rsh8x64 <t> n (Const64 <typ.UInt64> [ 7])) (Const64 <typ.UInt64> [ 8-log2(c)]))) (Const64 <typ.UInt64> [log2(c)]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpRsh8x64)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v1 := b.NewValue0(v.Pos, OpRsh8Ux64, t)
		v2 := b.NewValue0(v.Pos, OpRsh8x64, t)
		v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v3.AuxInt = 7
		v2.AddArg2(n, v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 8 - log2(c)
		v1.AddArg2(v2, v4)
		v0.AddArg2(n, v1)
		v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v5.AuxInt = log2(c)
		v.AddArg2(v0, v5)
		return true
	}
	// match: (Div8 <t> x (Const8 [c]))
	// cond: smagicOK(8,c)
	// result: (Sub8 <t> (Rsh32x64 <t> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(smagic(8,c).m)]) (SignExt8to32 x)) (Const64 <typ.UInt64> [8+smagic(8,c).s])) (Rsh32x64 <t> (SignExt8to32 x) (Const64 <typ.UInt64> [31])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(smagicOK(8, c)) {
			break
		}
		v.reset(OpSub8)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(smagic(8, c).m)
		v3 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v3.AddArg(x)
		v1.AddArg2(v2, v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 8 + smagic(8, c).s
		v0.AddArg2(v1, v4)
		v5 := b.NewValue0(v.Pos, OpRsh32x64, t)
		v6 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v6.AddArg(x)
		v7 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v7.AuxInt = 31
		v5.AddArg2(v6, v7)
		v.AddArg2(v0, v5)
		return true
	}
	return false
}
func rewriteValuegeneric_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u (Const8 [c]) (Const8 [d]))
	// cond: d != 0
	// result: (Const8 [int64(int8(uint8(c)/uint8(d)))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = int64(int8(uint8(c) / uint8(d)))
		return true
	}
	// match: (Div8u n (Const8 [c]))
	// cond: isPowerOfTwo(c&0xff)
	// result: (Rsh8Ux64 n (Const64 <typ.UInt64> [log2(c&0xff)]))
	for {
		n := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xff)) {
			break
		}
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = log2(c & 0xff)
		v.AddArg2(n, v0)
		return true
	}
	// match: (Div8u x (Const8 [c]))
	// cond: umagicOK(8, c)
	// result: (Trunc32to8 (Rsh32Ux64 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(1<<8+umagic(8,c).m)]) (ZeroExt8to32 x)) (Const64 <typ.UInt64> [8+umagic(8,c).s])))
	for {
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(umagicOK(8, c)) {
			break
		}
		v.reset(OpTrunc32to8)
		v0 := b.NewValue0(v.Pos, OpRsh32Ux64, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
		v2.AuxInt = int64(1<<8 + umagic(8, c).m)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v3.AddArg(x)
		v1.AddArg2(v2, v3)
		v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v4.AuxInt = 8 + umagic(8, c).s
		v0.AddArg2(v1, v4)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Eq16 x x)
	// result: (ConstBool [1])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (Eq16 (Const16 <t> [c]) (Add16 (Const16 <t> [d]) x))
	// result: (Eq16 (Const16 <t> [int64(int16(c-d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpEq16)
				v0 := b.NewValue0(v.Pos, OpConst16, t)
				v0.AuxInt = int64(int16(c - d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Eq16 (Const16 [c]) (Const16 [d]))
	// result: (ConstBool [b2i(c == d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst16 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c == d)
			return true
		}
		break
	}
	// match: (Eq16 (Mod16u x (Const16 [c])) (Const16 [0]))
	// cond: x.Op != OpConst16 && udivisibleOK(16,c) && !hasSmallRotate(config)
	// result: (Eq32 (Mod32u <typ.UInt32> (ZeroExt16to32 <typ.UInt32> x) (Const32 <typ.UInt32> [c&0xffff])) (Const32 <typ.UInt32> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMod16u {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpConst16 {
				continue
			}
			c := v_0_1.AuxInt
			if v_1.Op != OpConst16 || v_1.AuxInt != 0 || !(x.Op != OpConst16 && udivisibleOK(16, c) && !hasSmallRotate(config)) {
				continue
			}
			v.reset(OpEq32)
			v0 := b.NewValue0(v.Pos, OpMod32u, typ.UInt32)
			v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
			v1.AddArg(x)
			v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
			v2.AuxInt = c & 0xffff
			v0.AddArg2(v1, v2)
			v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
			v3.AuxInt = 0
			v.AddArg2(v0, v3)
			return true
		}
		break
	}
	// match: (Eq16 (Mod16 x (Const16 [c])) (Const16 [0]))
	// cond: x.Op != OpConst16 && sdivisibleOK(16,c) && !hasSmallRotate(config)
	// result: (Eq32 (Mod32 <typ.Int32> (SignExt16to32 <typ.Int32> x) (Const32 <typ.Int32> [c])) (Const32 <typ.Int32> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMod16 {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpConst16 {
				continue
			}
			c := v_0_1.AuxInt
			if v_1.Op != OpConst16 || v_1.AuxInt != 0 || !(x.Op != OpConst16 && sdivisibleOK(16, c) && !hasSmallRotate(config)) {
				continue
			}
			v.reset(OpEq32)
			v0 := b.NewValue0(v.Pos, OpMod32, typ.Int32)
			v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
			v1.AddArg(x)
			v2 := b.NewValue0(v.Pos, OpConst32, typ.Int32)
			v2.AuxInt = c
			v0.AddArg2(v1, v2)
			v3 := b.NewValue0(v.Pos, OpConst32, typ.Int32)
			v3.AuxInt = 0
			v.AddArg2(v0, v3)
			return true
		}
		break
	}
	// match: (Eq16 x (Mul16 (Const16 [c]) (Trunc64to16 (Rsh64Ux64 mul:(Mul64 (Const64 [m]) (ZeroExt16to64 x)) (Const64 [s]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<16+umagic(16,c).m) && s == 16+umagic(16,c).s && x.Op != OpConst16 && udivisibleOK(16,c)
	// result: (Leq16U (RotateLeft16 <typ.UInt16> (Mul16 <typ.UInt16> (Const16 <typ.UInt16> [int64(int16(udivisible(16,c).m))]) x) (Const16 <typ.UInt16> [int64(16-udivisible(16,c).k)]) ) (Const16 <typ.UInt16> [int64(int16(udivisible(16,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpTrunc64to16 {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh64Ux64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				mul := v_1_1_0.Args[0]
				if mul.Op != OpMul64 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst64 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpZeroExt16to64 || x != mul_1.Args[0] {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<16+umagic(16, c).m) && s == 16+umagic(16, c).s && x.Op != OpConst16 && udivisibleOK(16, c)) {
						continue
					}
					v.reset(OpLeq16U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft16, typ.UInt16)
					v1 := b.NewValue0(v.Pos, OpMul16, typ.UInt16)
					v2 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v2.AuxInt = int64(int16(udivisible(16, c).m))
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v3.AuxInt = int64(16 - udivisible(16, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v4.AuxInt = int64(int16(udivisible(16, c).max))
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq16 x (Mul16 (Const16 [c]) (Trunc32to16 (Rsh32Ux64 mul:(Mul32 (Const32 [m]) (ZeroExt16to32 x)) (Const64 [s]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<15+umagic(16,c).m/2) && s == 16+umagic(16,c).s-1 && x.Op != OpConst16 && udivisibleOK(16,c)
	// result: (Leq16U (RotateLeft16 <typ.UInt16> (Mul16 <typ.UInt16> (Const16 <typ.UInt16> [int64(int16(udivisible(16,c).m))]) x) (Const16 <typ.UInt16> [int64(16-udivisible(16,c).k)]) ) (Const16 <typ.UInt16> [int64(int16(udivisible(16,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpTrunc32to16 {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh32Ux64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				mul := v_1_1_0.Args[0]
				if mul.Op != OpMul32 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst32 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpZeroExt16to32 || x != mul_1.Args[0] {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<15+umagic(16, c).m/2) && s == 16+umagic(16, c).s-1 && x.Op != OpConst16 && udivisibleOK(16, c)) {
						continue
					}
					v.reset(OpLeq16U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft16, typ.UInt16)
					v1 := b.NewValue0(v.Pos, OpMul16, typ.UInt16)
					v2 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v2.AuxInt = int64(int16(udivisible(16, c).m))
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v3.AuxInt = int64(16 - udivisible(16, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v4.AuxInt = int64(int16(udivisible(16, c).max))
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq16 x (Mul16 (Const16 [c]) (Trunc32to16 (Rsh32Ux64 mul:(Mul32 (Const32 [m]) (Rsh32Ux64 (ZeroExt16to32 x) (Const64 [1]))) (Const64 [s]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<15+(umagic(16,c).m+1)/2) && s == 16+umagic(16,c).s-2 && x.Op != OpConst16 && udivisibleOK(16,c)
	// result: (Leq16U (RotateLeft16 <typ.UInt16> (Mul16 <typ.UInt16> (Const16 <typ.UInt16> [int64(int16(udivisible(16,c).m))]) x) (Const16 <typ.UInt16> [int64(16-udivisible(16,c).k)]) ) (Const16 <typ.UInt16> [int64(int16(udivisible(16,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpTrunc32to16 {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh32Ux64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				mul := v_1_1_0.Args[0]
				if mul.Op != OpMul32 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst32 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpRsh32Ux64 {
						continue
					}
					_ = mul_1.Args[1]
					mul_1_0 := mul_1.Args[0]
					if mul_1_0.Op != OpZeroExt16to32 || x != mul_1_0.Args[0] {
						continue
					}
					mul_1_1 := mul_1.Args[1]
					if mul_1_1.Op != OpConst64 || mul_1_1.AuxInt != 1 {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<15+(umagic(16, c).m+1)/2) && s == 16+umagic(16, c).s-2 && x.Op != OpConst16 && udivisibleOK(16, c)) {
						continue
					}
					v.reset(OpLeq16U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft16, typ.UInt16)
					v1 := b.NewValue0(v.Pos, OpMul16, typ.UInt16)
					v2 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v2.AuxInt = int64(int16(udivisible(16, c).m))
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v3.AuxInt = int64(16 - udivisible(16, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v4.AuxInt = int64(int16(udivisible(16, c).max))
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq16 x (Mul16 (Const16 [c]) (Trunc32to16 (Rsh32Ux64 (Avg32u (Lsh32x64 (ZeroExt16to32 x) (Const64 [16])) mul:(Mul32 (Const32 [m]) (ZeroExt16to32 x))) (Const64 [s]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(umagic(16,c).m) && s == 16+umagic(16,c).s-1 && x.Op != OpConst16 && udivisibleOK(16,c)
	// result: (Leq16U (RotateLeft16 <typ.UInt16> (Mul16 <typ.UInt16> (Const16 <typ.UInt16> [int64(int16(udivisible(16,c).m))]) x) (Const16 <typ.UInt16> [int64(16-udivisible(16,c).k)]) ) (Const16 <typ.UInt16> [int64(int16(udivisible(16,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpTrunc32to16 {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh32Ux64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAvg32u {
					continue
				}
				_ = v_1_1_0_0.Args[1]
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpLsh32x64 {
					continue
				}
				_ = v_1_1_0_0_0.Args[1]
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpZeroExt16to32 || x != v_1_1_0_0_0_0.Args[0] {
					continue
				}
				v_1_1_0_0_0_1 := v_1_1_0_0_0.Args[1]
				if v_1_1_0_0_0_1.Op != OpConst64 || v_1_1_0_0_0_1.AuxInt != 16 {
					continue
				}
				mul := v_1_1_0_0.Args[1]
				if mul.Op != OpMul32 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst32 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpZeroExt16to32 || x != mul_1.Args[0] {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(umagic(16, c).m) && s == 16+umagic(16, c).s-1 && x.Op != OpConst16 && udivisibleOK(16, c)) {
						continue
					}
					v.reset(OpLeq16U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft16, typ.UInt16)
					v1 := b.NewValue0(v.Pos, OpMul16, typ.UInt16)
					v2 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v2.AuxInt = int64(int16(udivisible(16, c).m))
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v3.AuxInt = int64(16 - udivisible(16, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v4.AuxInt = int64(int16(udivisible(16, c).max))
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq16 x (Mul16 (Const16 [c]) (Sub16 (Rsh32x64 mul:(Mul32 (Const32 [m]) (SignExt16to32 x)) (Const64 [s])) (Rsh32x64 (SignExt16to32 x) (Const64 [31]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(smagic(16,c).m) && s == 16+smagic(16,c).s && x.Op != OpConst16 && sdivisibleOK(16,c)
	// result: (Leq16U (RotateLeft16 <typ.UInt16> (Add16 <typ.UInt16> (Mul16 <typ.UInt16> (Const16 <typ.UInt16> [int64(int16(sdivisible(16,c).m))]) x) (Const16 <typ.UInt16> [int64(int16(sdivisible(16,c).a))]) ) (Const16 <typ.UInt16> [int64(16-sdivisible(16,c).k)]) ) (Const16 <typ.UInt16> [int64(int16(sdivisible(16,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpSub16 {
					continue
				}
				_ = v_1_1.Args[1]
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh32x64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				mul := v_1_1_0.Args[0]
				if mul.Op != OpMul32 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst32 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpSignExt16to32 || x != mul_1.Args[0] {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					v_1_1_1 := v_1_1.Args[1]
					if v_1_1_1.Op != OpRsh32x64 {
						continue
					}
					_ = v_1_1_1.Args[1]
					v_1_1_1_0 := v_1_1_1.Args[0]
					if v_1_1_1_0.Op != OpSignExt16to32 || x != v_1_1_1_0.Args[0] {
						continue
					}
					v_1_1_1_1 := v_1_1_1.Args[1]
					if v_1_1_1_1.Op != OpConst64 || v_1_1_1_1.AuxInt != 31 || !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(smagic(16, c).m) && s == 16+smagic(16, c).s && x.Op != OpConst16 && sdivisibleOK(16, c)) {
						continue
					}
					v.reset(OpLeq16U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft16, typ.UInt16)
					v1 := b.NewValue0(v.Pos, OpAdd16, typ.UInt16)
					v2 := b.NewValue0(v.Pos, OpMul16, typ.UInt16)
					v3 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v3.AuxInt = int64(int16(sdivisible(16, c).m))
					v2.AddArg2(v3, x)
					v4 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v4.AuxInt = int64(int16(sdivisible(16, c).a))
					v1.AddArg2(v2, v4)
					v5 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v5.AuxInt = int64(16 - sdivisible(16, c).k)
					v0.AddArg2(v1, v5)
					v6 := b.NewValue0(v.Pos, OpConst16, typ.UInt16)
					v6.AuxInt = int64(int16(sdivisible(16, c).max))
					v.AddArg2(v0, v6)
					return true
				}
			}
		}
		break
	}
	// match: (Eq16 n (Lsh16x64 (Rsh16x64 (Add16 <t> n (Rsh16Ux64 <t> (Rsh16x64 <t> n (Const64 <typ.UInt64> [15])) (Const64 <typ.UInt64> [kbar]))) (Const64 <typ.UInt64> [k])) (Const64 <typ.UInt64> [k])) )
	// cond: k > 0 && k < 15 && kbar == 16 - k
	// result: (Eq16 (And16 <t> n (Const16 <t> [int64(1<<uint(k)-1)])) (Const16 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpLsh16x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpRsh16x64 {
				continue
			}
			_ = v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAdd16 {
				continue
			}
			t := v_1_0_0.Type
			_ = v_1_0_0.Args[1]
			v_1_0_0_0 := v_1_0_0.Args[0]
			v_1_0_0_1 := v_1_0_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0_0_0, v_1_0_0_1 = _i1+1, v_1_0_0_1, v_1_0_0_0 {
				if n != v_1_0_0_0 || v_1_0_0_1.Op != OpRsh16Ux64 || v_1_0_0_1.Type != t {
					continue
				}
				_ = v_1_0_0_1.Args[1]
				v_1_0_0_1_0 := v_1_0_0_1.Args[0]
				if v_1_0_0_1_0.Op != OpRsh16x64 || v_1_0_0_1_0.Type != t {
					continue
				}
				_ = v_1_0_0_1_0.Args[1]
				if n != v_1_0_0_1_0.Args[0] {
					continue
				}
				v_1_0_0_1_0_1 := v_1_0_0_1_0.Args[1]
				if v_1_0_0_1_0_1.Op != OpConst64 || v_1_0_0_1_0_1.Type != typ.UInt64 || v_1_0_0_1_0_1.AuxInt != 15 {
					continue
				}
				v_1_0_0_1_1 := v_1_0_0_1.Args[1]
				if v_1_0_0_1_1.Op != OpConst64 || v_1_0_0_1_1.Type != typ.UInt64 {
					continue
				}
				kbar := v_1_0_0_1_1.AuxInt
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpConst64 || v_1_0_1.Type != typ.UInt64 {
					continue
				}
				k := v_1_0_1.AuxInt
				v_1_1 := v_1.Args[1]
				if v_1_1.Op != OpConst64 || v_1_1.Type != typ.UInt64 || v_1_1.AuxInt != k || !(k > 0 && k < 15 && kbar == 16-k) {
					continue
				}
				v.reset(OpEq16)
				v0 := b.NewValue0(v.Pos, OpAnd16, t)
				v1 := b.NewValue0(v.Pos, OpConst16, t)
				v1.AuxInt = int64(1<<uint(k) - 1)
				v0.AddArg2(n, v1)
				v2 := b.NewValue0(v.Pos, OpConst16, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	// match: (Eq16 s:(Sub16 x y) (Const16 [0]))
	// cond: s.Uses == 1
	// result: (Eq16 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s := v_0
			if s.Op != OpSub16 {
				continue
			}
			y := s.Args[1]
			x := s.Args[0]
			if v_1.Op != OpConst16 || v_1.AuxInt != 0 || !(s.Uses == 1) {
				continue
			}
			v.reset(OpEq16)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (Eq16 (And16 <t> x (Const16 <t> [y])) (Const16 <t> [y]))
	// cond: isPowerOfTwo(y)
	// result: (Neq16 (And16 <t> x (Const16 <t> [y])) (Const16 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd16 {
				continue
			}
			t := v_0.Type
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst16 || v_0_1.Type != t {
					continue
				}
				y := v_0_1.AuxInt
				if v_1.Op != OpConst16 || v_1.Type != t || v_1.AuxInt != y || !(isPowerOfTwo(y)) {
					continue
				}
				v.reset(OpNeq16)
				v0 := b.NewValue0(v.Pos, OpAnd16, t)
				v1 := b.NewValue0(v.Pos, OpConst16, t)
				v1.AuxInt = y
				v0.AddArg2(x, v1)
				v2 := b.NewValue0(v.Pos, OpConst16, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq32 x x)
	// result: (ConstBool [1])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (Eq32 (Const32 <t> [c]) (Add32 (Const32 <t> [d]) x))
	// result: (Eq32 (Const32 <t> [int64(int32(c-d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpEq32)
				v0 := b.NewValue0(v.Pos, OpConst32, t)
				v0.AuxInt = int64(int32(c - d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Eq32 (Const32 [c]) (Const32 [d]))
	// result: (ConstBool [b2i(c == d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c == d)
			return true
		}
		break
	}
	// match: (Eq32 x (Mul32 (Const32 [c]) (Rsh32Ux64 mul:(Hmul32u (Const32 [m]) x) (Const64 [s])) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(int32(1<<31+umagic(32,c).m/2)) && s == umagic(32,c).s-1 && x.Op != OpConst32 && udivisibleOK(32,c)
	// result: (Leq32U (RotateLeft32 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).m))]) x) (Const32 <typ.UInt32> [int64(32-udivisible(32,c).k)]) ) (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpRsh32Ux64 {
					continue
				}
				_ = v_1_1.Args[1]
				mul := v_1_1.Args[0]
				if mul.Op != OpHmul32u {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst32 {
						continue
					}
					m := mul_0.AuxInt
					if x != mul_1 {
						continue
					}
					v_1_1_1 := v_1_1.Args[1]
					if v_1_1_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(int32(1<<31+umagic(32, c).m/2)) && s == umagic(32, c).s-1 && x.Op != OpConst32 && udivisibleOK(32, c)) {
						continue
					}
					v.reset(OpLeq32U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft32, typ.UInt32)
					v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
					v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v2.AuxInt = int64(int32(udivisible(32, c).m))
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v3.AuxInt = int64(32 - udivisible(32, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v4.AuxInt = int64(int32(udivisible(32, c).max))
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq32 x (Mul32 (Const32 [c]) (Rsh32Ux64 mul:(Hmul32u (Const32 <typ.UInt32> [m]) (Rsh32Ux64 x (Const64 [1]))) (Const64 [s])) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(int32(1<<31+(umagic(32,c).m+1)/2)) && s == umagic(32,c).s-2 && x.Op != OpConst32 && udivisibleOK(32,c)
	// result: (Leq32U (RotateLeft32 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).m))]) x) (Const32 <typ.UInt32> [int64(32-udivisible(32,c).k)]) ) (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpRsh32Ux64 {
					continue
				}
				_ = v_1_1.Args[1]
				mul := v_1_1.Args[0]
				if mul.Op != OpHmul32u {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst32 || mul_0.Type != typ.UInt32 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpRsh32Ux64 {
						continue
					}
					_ = mul_1.Args[1]
					if x != mul_1.Args[0] {
						continue
					}
					mul_1_1 := mul_1.Args[1]
					if mul_1_1.Op != OpConst64 || mul_1_1.AuxInt != 1 {
						continue
					}
					v_1_1_1 := v_1_1.Args[1]
					if v_1_1_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(int32(1<<31+(umagic(32, c).m+1)/2)) && s == umagic(32, c).s-2 && x.Op != OpConst32 && udivisibleOK(32, c)) {
						continue
					}
					v.reset(OpLeq32U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft32, typ.UInt32)
					v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
					v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v2.AuxInt = int64(int32(udivisible(32, c).m))
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v3.AuxInt = int64(32 - udivisible(32, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v4.AuxInt = int64(int32(udivisible(32, c).max))
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq32 x (Mul32 (Const32 [c]) (Rsh32Ux64 (Avg32u x mul:(Hmul32u (Const32 [m]) x)) (Const64 [s])) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(int32(umagic(32,c).m)) && s == umagic(32,c).s-1 && x.Op != OpConst32 && udivisibleOK(32,c)
	// result: (Leq32U (RotateLeft32 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).m))]) x) (Const32 <typ.UInt32> [int64(32-udivisible(32,c).k)]) ) (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpRsh32Ux64 {
					continue
				}
				_ = v_1_1.Args[1]
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAvg32u {
					continue
				}
				_ = v_1_1_0.Args[1]
				if x != v_1_1_0.Args[0] {
					continue
				}
				mul := v_1_1_0.Args[1]
				if mul.Op != OpHmul32u {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst32 {
						continue
					}
					m := mul_0.AuxInt
					if x != mul_1 {
						continue
					}
					v_1_1_1 := v_1_1.Args[1]
					if v_1_1_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(int32(umagic(32, c).m)) && s == umagic(32, c).s-1 && x.Op != OpConst32 && udivisibleOK(32, c)) {
						continue
					}
					v.reset(OpLeq32U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft32, typ.UInt32)
					v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
					v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v2.AuxInt = int64(int32(udivisible(32, c).m))
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v3.AuxInt = int64(32 - udivisible(32, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v4.AuxInt = int64(int32(udivisible(32, c).max))
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq32 x (Mul32 (Const32 [c]) (Trunc64to32 (Rsh64Ux64 mul:(Mul64 (Const64 [m]) (ZeroExt32to64 x)) (Const64 [s]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<31+umagic(32,c).m/2) && s == 32+umagic(32,c).s-1 && x.Op != OpConst32 && udivisibleOK(32,c)
	// result: (Leq32U (RotateLeft32 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).m))]) x) (Const32 <typ.UInt32> [int64(32-udivisible(32,c).k)]) ) (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpTrunc64to32 {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh64Ux64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				mul := v_1_1_0.Args[0]
				if mul.Op != OpMul64 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst64 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpZeroExt32to64 || x != mul_1.Args[0] {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<31+umagic(32, c).m/2) && s == 32+umagic(32, c).s-1 && x.Op != OpConst32 && udivisibleOK(32, c)) {
						continue
					}
					v.reset(OpLeq32U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft32, typ.UInt32)
					v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
					v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v2.AuxInt = int64(int32(udivisible(32, c).m))
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v3.AuxInt = int64(32 - udivisible(32, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v4.AuxInt = int64(int32(udivisible(32, c).max))
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq32 x (Mul32 (Const32 [c]) (Trunc64to32 (Rsh64Ux64 mul:(Mul64 (Const64 [m]) (Rsh64Ux64 (ZeroExt32to64 x) (Const64 [1]))) (Const64 [s]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<31+(umagic(32,c).m+1)/2) && s == 32+umagic(32,c).s-2 && x.Op != OpConst32 && udivisibleOK(32,c)
	// result: (Leq32U (RotateLeft32 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).m))]) x) (Const32 <typ.UInt32> [int64(32-udivisible(32,c).k)]) ) (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpTrunc64to32 {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh64Ux64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				mul := v_1_1_0.Args[0]
				if mul.Op != OpMul64 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst64 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpRsh64Ux64 {
						continue
					}
					_ = mul_1.Args[1]
					mul_1_0 := mul_1.Args[0]
					if mul_1_0.Op != OpZeroExt32to64 || x != mul_1_0.Args[0] {
						continue
					}
					mul_1_1 := mul_1.Args[1]
					if mul_1_1.Op != OpConst64 || mul_1_1.AuxInt != 1 {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<31+(umagic(32, c).m+1)/2) && s == 32+umagic(32, c).s-2 && x.Op != OpConst32 && udivisibleOK(32, c)) {
						continue
					}
					v.reset(OpLeq32U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft32, typ.UInt32)
					v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
					v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v2.AuxInt = int64(int32(udivisible(32, c).m))
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v3.AuxInt = int64(32 - udivisible(32, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v4.AuxInt = int64(int32(udivisible(32, c).max))
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq32 x (Mul32 (Const32 [c]) (Trunc64to32 (Rsh64Ux64 (Avg64u (Lsh64x64 (ZeroExt32to64 x) (Const64 [32])) mul:(Mul64 (Const64 [m]) (ZeroExt32to64 x))) (Const64 [s]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(umagic(32,c).m) && s == 32+umagic(32,c).s-1 && x.Op != OpConst32 && udivisibleOK(32,c)
	// result: (Leq32U (RotateLeft32 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).m))]) x) (Const32 <typ.UInt32> [int64(32-udivisible(32,c).k)]) ) (Const32 <typ.UInt32> [int64(int32(udivisible(32,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpTrunc64to32 {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh64Ux64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAvg64u {
					continue
				}
				_ = v_1_1_0_0.Args[1]
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				if v_1_1_0_0_0.Op != OpLsh64x64 {
					continue
				}
				_ = v_1_1_0_0_0.Args[1]
				v_1_1_0_0_0_0 := v_1_1_0_0_0.Args[0]
				if v_1_1_0_0_0_0.Op != OpZeroExt32to64 || x != v_1_1_0_0_0_0.Args[0] {
					continue
				}
				v_1_1_0_0_0_1 := v_1_1_0_0_0.Args[1]
				if v_1_1_0_0_0_1.Op != OpConst64 || v_1_1_0_0_0_1.AuxInt != 32 {
					continue
				}
				mul := v_1_1_0_0.Args[1]
				if mul.Op != OpMul64 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst64 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpZeroExt32to64 || x != mul_1.Args[0] {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(umagic(32, c).m) && s == 32+umagic(32, c).s-1 && x.Op != OpConst32 && udivisibleOK(32, c)) {
						continue
					}
					v.reset(OpLeq32U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft32, typ.UInt32)
					v1 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
					v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v2.AuxInt = int64(int32(udivisible(32, c).m))
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v3.AuxInt = int64(32 - udivisible(32, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v4.AuxInt = int64(int32(udivisible(32, c).max))
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq32 x (Mul32 (Const32 [c]) (Sub32 (Rsh64x64 mul:(Mul64 (Const64 [m]) (SignExt32to64 x)) (Const64 [s])) (Rsh64x64 (SignExt32to64 x) (Const64 [63]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(smagic(32,c).m) && s == 32+smagic(32,c).s && x.Op != OpConst32 && sdivisibleOK(32,c)
	// result: (Leq32U (RotateLeft32 <typ.UInt32> (Add32 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(sdivisible(32,c).m))]) x) (Const32 <typ.UInt32> [int64(int32(sdivisible(32,c).a))]) ) (Const32 <typ.UInt32> [int64(32-sdivisible(32,c).k)]) ) (Const32 <typ.UInt32> [int64(int32(sdivisible(32,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpSub32 {
					continue
				}
				_ = v_1_1.Args[1]
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh64x64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				mul := v_1_1_0.Args[0]
				if mul.Op != OpMul64 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst64 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpSignExt32to64 || x != mul_1.Args[0] {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					v_1_1_1 := v_1_1.Args[1]
					if v_1_1_1.Op != OpRsh64x64 {
						continue
					}
					_ = v_1_1_1.Args[1]
					v_1_1_1_0 := v_1_1_1.Args[0]
					if v_1_1_1_0.Op != OpSignExt32to64 || x != v_1_1_1_0.Args[0] {
						continue
					}
					v_1_1_1_1 := v_1_1_1.Args[1]
					if v_1_1_1_1.Op != OpConst64 || v_1_1_1_1.AuxInt != 63 || !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(smagic(32, c).m) && s == 32+smagic(32, c).s && x.Op != OpConst32 && sdivisibleOK(32, c)) {
						continue
					}
					v.reset(OpLeq32U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft32, typ.UInt32)
					v1 := b.NewValue0(v.Pos, OpAdd32, typ.UInt32)
					v2 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
					v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v3.AuxInt = int64(int32(sdivisible(32, c).m))
					v2.AddArg2(v3, x)
					v4 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v4.AuxInt = int64(int32(sdivisible(32, c).a))
					v1.AddArg2(v2, v4)
					v5 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v5.AuxInt = int64(32 - sdivisible(32, c).k)
					v0.AddArg2(v1, v5)
					v6 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v6.AuxInt = int64(int32(sdivisible(32, c).max))
					v.AddArg2(v0, v6)
					return true
				}
			}
		}
		break
	}
	// match: (Eq32 x (Mul32 (Const32 [c]) (Sub32 (Rsh32x64 mul:(Hmul32 (Const32 [m]) x) (Const64 [s])) (Rsh32x64 x (Const64 [31]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(int32(smagic(32,c).m/2)) && s == smagic(32,c).s-1 && x.Op != OpConst32 && sdivisibleOK(32,c)
	// result: (Leq32U (RotateLeft32 <typ.UInt32> (Add32 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(sdivisible(32,c).m))]) x) (Const32 <typ.UInt32> [int64(int32(sdivisible(32,c).a))]) ) (Const32 <typ.UInt32> [int64(32-sdivisible(32,c).k)]) ) (Const32 <typ.UInt32> [int64(int32(sdivisible(32,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpSub32 {
					continue
				}
				_ = v_1_1.Args[1]
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh32x64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				mul := v_1_1_0.Args[0]
				if mul.Op != OpHmul32 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst32 {
						continue
					}
					m := mul_0.AuxInt
					if x != mul_1 {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					v_1_1_1 := v_1_1.Args[1]
					if v_1_1_1.Op != OpRsh32x64 {
						continue
					}
					_ = v_1_1_1.Args[1]
					if x != v_1_1_1.Args[0] {
						continue
					}
					v_1_1_1_1 := v_1_1_1.Args[1]
					if v_1_1_1_1.Op != OpConst64 || v_1_1_1_1.AuxInt != 31 || !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(int32(smagic(32, c).m/2)) && s == smagic(32, c).s-1 && x.Op != OpConst32 && sdivisibleOK(32, c)) {
						continue
					}
					v.reset(OpLeq32U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft32, typ.UInt32)
					v1 := b.NewValue0(v.Pos, OpAdd32, typ.UInt32)
					v2 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
					v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v3.AuxInt = int64(int32(sdivisible(32, c).m))
					v2.AddArg2(v3, x)
					v4 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v4.AuxInt = int64(int32(sdivisible(32, c).a))
					v1.AddArg2(v2, v4)
					v5 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v5.AuxInt = int64(32 - sdivisible(32, c).k)
					v0.AddArg2(v1, v5)
					v6 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
					v6.AuxInt = int64(int32(sdivisible(32, c).max))
					v.AddArg2(v0, v6)
					return true
				}
			}
		}
		break
	}
	// match: (Eq32 x (Mul32 (Const32 [c]) (Sub32 (Rsh32x64 (Add32 mul:(Hmul32 (Const32 [m]) x) x) (Const64 [s])) (Rsh32x64 x (Const64 [31]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(int32(smagic(32,c).m)) && s == smagic(32,c).s && x.Op != OpConst32 && sdivisibleOK(32,c)
	// result: (Leq32U (RotateLeft32 <typ.UInt32> (Add32 <typ.UInt32> (Mul32 <typ.UInt32> (Const32 <typ.UInt32> [int64(int32(sdivisible(32,c).m))]) x) (Const32 <typ.UInt32> [int64(int32(sdivisible(32,c).a))]) ) (Const32 <typ.UInt32> [int64(32-sdivisible(32,c).k)]) ) (Const32 <typ.UInt32> [int64(int32(sdivisible(32,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpSub32 {
					continue
				}
				_ = v_1_1.Args[1]
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh32x64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAdd32 {
					continue
				}
				_ = v_1_1_0_0.Args[1]
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				v_1_1_0_0_1 := v_1_1_0_0.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, v_1_1_0_0_0, v_1_1_0_0_1 = _i2+1, v_1_1_0_0_1, v_1_1_0_0_0 {
					mul := v_1_1_0_0_0
					if mul.Op != OpHmul32 {
						continue
					}
					_ = mul.Args[1]
					mul_0 := mul.Args[0]
					mul_1 := mul.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, mul_0, mul_1 = _i3+1, mul_1, mul_0 {
						if mul_0.Op != OpConst32 {
							continue
						}
						m := mul_0.AuxInt
						if x != mul_1 || x != v_1_1_0_0_1 {
							continue
						}
						v_1_1_0_1 := v_1_1_0.Args[1]
						if v_1_1_0_1.Op != OpConst64 {
							continue
						}
						s := v_1_1_0_1.AuxInt
						v_1_1_1 := v_1_1.Args[1]
						if v_1_1_1.Op != OpRsh32x64 {
							continue
						}
						_ = v_1_1_1.Args[1]
						if x != v_1_1_1.Args[0] {
							continue
						}
						v_1_1_1_1 := v_1_1_1.Args[1]
						if v_1_1_1_1.Op != OpConst64 || v_1_1_1_1.AuxInt != 31 || !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(int32(smagic(32, c).m)) && s == smagic(32, c).s && x.Op != OpConst32 && sdivisibleOK(32, c)) {
							continue
						}
						v.reset(OpLeq32U)
						v0 := b.NewValue0(v.Pos, OpRotateLeft32, typ.UInt32)
						v1 := b.NewValue0(v.Pos, OpAdd32, typ.UInt32)
						v2 := b.NewValue0(v.Pos, OpMul32, typ.UInt32)
						v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
						v3.AuxInt = int64(int32(sdivisible(32, c).m))
						v2.AddArg2(v3, x)
						v4 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
						v4.AuxInt = int64(int32(sdivisible(32, c).a))
						v1.AddArg2(v2, v4)
						v5 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
						v5.AuxInt = int64(32 - sdivisible(32, c).k)
						v0.AddArg2(v1, v5)
						v6 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
						v6.AuxInt = int64(int32(sdivisible(32, c).max))
						v.AddArg2(v0, v6)
						return true
					}
				}
			}
		}
		break
	}
	// match: (Eq32 n (Lsh32x64 (Rsh32x64 (Add32 <t> n (Rsh32Ux64 <t> (Rsh32x64 <t> n (Const64 <typ.UInt64> [31])) (Const64 <typ.UInt64> [kbar]))) (Const64 <typ.UInt64> [k])) (Const64 <typ.UInt64> [k])) )
	// cond: k > 0 && k < 31 && kbar == 32 - k
	// result: (Eq32 (And32 <t> n (Const32 <t> [int64(1<<uint(k)-1)])) (Const32 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpLsh32x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpRsh32x64 {
				continue
			}
			_ = v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAdd32 {
				continue
			}
			t := v_1_0_0.Type
			_ = v_1_0_0.Args[1]
			v_1_0_0_0 := v_1_0_0.Args[0]
			v_1_0_0_1 := v_1_0_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0_0_0, v_1_0_0_1 = _i1+1, v_1_0_0_1, v_1_0_0_0 {
				if n != v_1_0_0_0 || v_1_0_0_1.Op != OpRsh32Ux64 || v_1_0_0_1.Type != t {
					continue
				}
				_ = v_1_0_0_1.Args[1]
				v_1_0_0_1_0 := v_1_0_0_1.Args[0]
				if v_1_0_0_1_0.Op != OpRsh32x64 || v_1_0_0_1_0.Type != t {
					continue
				}
				_ = v_1_0_0_1_0.Args[1]
				if n != v_1_0_0_1_0.Args[0] {
					continue
				}
				v_1_0_0_1_0_1 := v_1_0_0_1_0.Args[1]
				if v_1_0_0_1_0_1.Op != OpConst64 || v_1_0_0_1_0_1.Type != typ.UInt64 || v_1_0_0_1_0_1.AuxInt != 31 {
					continue
				}
				v_1_0_0_1_1 := v_1_0_0_1.Args[1]
				if v_1_0_0_1_1.Op != OpConst64 || v_1_0_0_1_1.Type != typ.UInt64 {
					continue
				}
				kbar := v_1_0_0_1_1.AuxInt
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpConst64 || v_1_0_1.Type != typ.UInt64 {
					continue
				}
				k := v_1_0_1.AuxInt
				v_1_1 := v_1.Args[1]
				if v_1_1.Op != OpConst64 || v_1_1.Type != typ.UInt64 || v_1_1.AuxInt != k || !(k > 0 && k < 31 && kbar == 32-k) {
					continue
				}
				v.reset(OpEq32)
				v0 := b.NewValue0(v.Pos, OpAnd32, t)
				v1 := b.NewValue0(v.Pos, OpConst32, t)
				v1.AuxInt = int64(1<<uint(k) - 1)
				v0.AddArg2(n, v1)
				v2 := b.NewValue0(v.Pos, OpConst32, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	// match: (Eq32 s:(Sub32 x y) (Const32 [0]))
	// cond: s.Uses == 1
	// result: (Eq32 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s := v_0
			if s.Op != OpSub32 {
				continue
			}
			y := s.Args[1]
			x := s.Args[0]
			if v_1.Op != OpConst32 || v_1.AuxInt != 0 || !(s.Uses == 1) {
				continue
			}
			v.reset(OpEq32)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (Eq32 (And32 <t> x (Const32 <t> [y])) (Const32 <t> [y]))
	// cond: isPowerOfTwo(y)
	// result: (Neq32 (And32 <t> x (Const32 <t> [y])) (Const32 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd32 {
				continue
			}
			t := v_0.Type
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst32 || v_0_1.Type != t {
					continue
				}
				y := v_0_1.AuxInt
				if v_1.Op != OpConst32 || v_1.Type != t || v_1.AuxInt != y || !(isPowerOfTwo(y)) {
					continue
				}
				v.reset(OpNeq32)
				v0 := b.NewValue0(v.Pos, OpAnd32, t)
				v1 := b.NewValue0(v.Pos, OpConst32, t)
				v1.AuxInt = y
				v0.AddArg2(x, v1)
				v2 := b.NewValue0(v.Pos, OpConst32, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpEq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Eq32F (Const32F [c]) (Const32F [d]))
	// result: (ConstBool [b2i(auxTo32F(c) == auxTo32F(d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32F {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32F {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(auxTo32F(c) == auxTo32F(d))
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpEq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq64 x x)
	// result: (ConstBool [1])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (Eq64 (Const64 <t> [c]) (Add64 (Const64 <t> [d]) x))
	// result: (Eq64 (Const64 <t> [c-d]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpEq64)
				v0 := b.NewValue0(v.Pos, OpConst64, t)
				v0.AuxInt = c - d
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Eq64 (Const64 [c]) (Const64 [d]))
	// result: (ConstBool [b2i(c == d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c == d)
			return true
		}
		break
	}
	// match: (Eq64 x (Mul64 (Const64 [c]) (Rsh64Ux64 mul:(Hmul64u (Const64 [m]) x) (Const64 [s])) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<63+umagic(64,c).m/2) && s == umagic(64,c).s-1 && x.Op != OpConst64 && udivisibleOK(64,c)
	// result: (Leq64U (RotateLeft64 <typ.UInt64> (Mul64 <typ.UInt64> (Const64 <typ.UInt64> [int64(udivisible(64,c).m)]) x) (Const64 <typ.UInt64> [int64(64-udivisible(64,c).k)]) ) (Const64 <typ.UInt64> [int64(udivisible(64,c).max)]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpRsh64Ux64 {
					continue
				}
				_ = v_1_1.Args[1]
				mul := v_1_1.Args[0]
				if mul.Op != OpHmul64u {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst64 {
						continue
					}
					m := mul_0.AuxInt
					if x != mul_1 {
						continue
					}
					v_1_1_1 := v_1_1.Args[1]
					if v_1_1_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<63+umagic(64, c).m/2) && s == umagic(64, c).s-1 && x.Op != OpConst64 && udivisibleOK(64, c)) {
						continue
					}
					v.reset(OpLeq64U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft64, typ.UInt64)
					v1 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
					v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v2.AuxInt = int64(udivisible(64, c).m)
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v3.AuxInt = int64(64 - udivisible(64, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v4.AuxInt = int64(udivisible(64, c).max)
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq64 x (Mul64 (Const64 [c]) (Rsh64Ux64 mul:(Hmul64u (Const64 [m]) (Rsh64Ux64 x (Const64 [1]))) (Const64 [s])) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<63+(umagic(64,c).m+1)/2) && s == umagic(64,c).s-2 && x.Op != OpConst64 && udivisibleOK(64,c)
	// result: (Leq64U (RotateLeft64 <typ.UInt64> (Mul64 <typ.UInt64> (Const64 <typ.UInt64> [int64(udivisible(64,c).m)]) x) (Const64 <typ.UInt64> [int64(64-udivisible(64,c).k)]) ) (Const64 <typ.UInt64> [int64(udivisible(64,c).max)]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpRsh64Ux64 {
					continue
				}
				_ = v_1_1.Args[1]
				mul := v_1_1.Args[0]
				if mul.Op != OpHmul64u {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst64 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpRsh64Ux64 {
						continue
					}
					_ = mul_1.Args[1]
					if x != mul_1.Args[0] {
						continue
					}
					mul_1_1 := mul_1.Args[1]
					if mul_1_1.Op != OpConst64 || mul_1_1.AuxInt != 1 {
						continue
					}
					v_1_1_1 := v_1_1.Args[1]
					if v_1_1_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<63+(umagic(64, c).m+1)/2) && s == umagic(64, c).s-2 && x.Op != OpConst64 && udivisibleOK(64, c)) {
						continue
					}
					v.reset(OpLeq64U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft64, typ.UInt64)
					v1 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
					v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v2.AuxInt = int64(udivisible(64, c).m)
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v3.AuxInt = int64(64 - udivisible(64, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v4.AuxInt = int64(udivisible(64, c).max)
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq64 x (Mul64 (Const64 [c]) (Rsh64Ux64 (Avg64u x mul:(Hmul64u (Const64 [m]) x)) (Const64 [s])) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(umagic(64,c).m) && s == umagic(64,c).s-1 && x.Op != OpConst64 && udivisibleOK(64,c)
	// result: (Leq64U (RotateLeft64 <typ.UInt64> (Mul64 <typ.UInt64> (Const64 <typ.UInt64> [int64(udivisible(64,c).m)]) x) (Const64 <typ.UInt64> [int64(64-udivisible(64,c).k)]) ) (Const64 <typ.UInt64> [int64(udivisible(64,c).max)]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpRsh64Ux64 {
					continue
				}
				_ = v_1_1.Args[1]
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpAvg64u {
					continue
				}
				_ = v_1_1_0.Args[1]
				if x != v_1_1_0.Args[0] {
					continue
				}
				mul := v_1_1_0.Args[1]
				if mul.Op != OpHmul64u {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst64 {
						continue
					}
					m := mul_0.AuxInt
					if x != mul_1 {
						continue
					}
					v_1_1_1 := v_1_1.Args[1]
					if v_1_1_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(umagic(64, c).m) && s == umagic(64, c).s-1 && x.Op != OpConst64 && udivisibleOK(64, c)) {
						continue
					}
					v.reset(OpLeq64U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft64, typ.UInt64)
					v1 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
					v2 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v2.AuxInt = int64(udivisible(64, c).m)
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v3.AuxInt = int64(64 - udivisible(64, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v4.AuxInt = int64(udivisible(64, c).max)
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq64 x (Mul64 (Const64 [c]) (Sub64 (Rsh64x64 mul:(Hmul64 (Const64 [m]) x) (Const64 [s])) (Rsh64x64 x (Const64 [63]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(smagic(64,c).m/2) && s == smagic(64,c).s-1 && x.Op != OpConst64 && sdivisibleOK(64,c)
	// result: (Leq64U (RotateLeft64 <typ.UInt64> (Add64 <typ.UInt64> (Mul64 <typ.UInt64> (Const64 <typ.UInt64> [int64(sdivisible(64,c).m)]) x) (Const64 <typ.UInt64> [int64(sdivisible(64,c).a)]) ) (Const64 <typ.UInt64> [int64(64-sdivisible(64,c).k)]) ) (Const64 <typ.UInt64> [int64(sdivisible(64,c).max)]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpSub64 {
					continue
				}
				_ = v_1_1.Args[1]
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh64x64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				mul := v_1_1_0.Args[0]
				if mul.Op != OpHmul64 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst64 {
						continue
					}
					m := mul_0.AuxInt
					if x != mul_1 {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					v_1_1_1 := v_1_1.Args[1]
					if v_1_1_1.Op != OpRsh64x64 {
						continue
					}
					_ = v_1_1_1.Args[1]
					if x != v_1_1_1.Args[0] {
						continue
					}
					v_1_1_1_1 := v_1_1_1.Args[1]
					if v_1_1_1_1.Op != OpConst64 || v_1_1_1_1.AuxInt != 63 || !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(smagic(64, c).m/2) && s == smagic(64, c).s-1 && x.Op != OpConst64 && sdivisibleOK(64, c)) {
						continue
					}
					v.reset(OpLeq64U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft64, typ.UInt64)
					v1 := b.NewValue0(v.Pos, OpAdd64, typ.UInt64)
					v2 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
					v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v3.AuxInt = int64(sdivisible(64, c).m)
					v2.AddArg2(v3, x)
					v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v4.AuxInt = int64(sdivisible(64, c).a)
					v1.AddArg2(v2, v4)
					v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v5.AuxInt = int64(64 - sdivisible(64, c).k)
					v0.AddArg2(v1, v5)
					v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
					v6.AuxInt = int64(sdivisible(64, c).max)
					v.AddArg2(v0, v6)
					return true
				}
			}
		}
		break
	}
	// match: (Eq64 x (Mul64 (Const64 [c]) (Sub64 (Rsh64x64 (Add64 mul:(Hmul64 (Const64 [m]) x) x) (Const64 [s])) (Rsh64x64 x (Const64 [63]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(smagic(64,c).m) && s == smagic(64,c).s && x.Op != OpConst64 && sdivisibleOK(64,c)
	// result: (Leq64U (RotateLeft64 <typ.UInt64> (Add64 <typ.UInt64> (Mul64 <typ.UInt64> (Const64 <typ.UInt64> [int64(sdivisible(64,c).m)]) x) (Const64 <typ.UInt64> [int64(sdivisible(64,c).a)]) ) (Const64 <typ.UInt64> [int64(64-sdivisible(64,c).k)]) ) (Const64 <typ.UInt64> [int64(sdivisible(64,c).max)]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpSub64 {
					continue
				}
				_ = v_1_1.Args[1]
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh64x64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				v_1_1_0_0 := v_1_1_0.Args[0]
				if v_1_1_0_0.Op != OpAdd64 {
					continue
				}
				_ = v_1_1_0_0.Args[1]
				v_1_1_0_0_0 := v_1_1_0_0.Args[0]
				v_1_1_0_0_1 := v_1_1_0_0.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, v_1_1_0_0_0, v_1_1_0_0_1 = _i2+1, v_1_1_0_0_1, v_1_1_0_0_0 {
					mul := v_1_1_0_0_0
					if mul.Op != OpHmul64 {
						continue
					}
					_ = mul.Args[1]
					mul_0 := mul.Args[0]
					mul_1 := mul.Args[1]
					for _i3 := 0; _i3 <= 1; _i3, mul_0, mul_1 = _i3+1, mul_1, mul_0 {
						if mul_0.Op != OpConst64 {
							continue
						}
						m := mul_0.AuxInt
						if x != mul_1 || x != v_1_1_0_0_1 {
							continue
						}
						v_1_1_0_1 := v_1_1_0.Args[1]
						if v_1_1_0_1.Op != OpConst64 {
							continue
						}
						s := v_1_1_0_1.AuxInt
						v_1_1_1 := v_1_1.Args[1]
						if v_1_1_1.Op != OpRsh64x64 {
							continue
						}
						_ = v_1_1_1.Args[1]
						if x != v_1_1_1.Args[0] {
							continue
						}
						v_1_1_1_1 := v_1_1_1.Args[1]
						if v_1_1_1_1.Op != OpConst64 || v_1_1_1_1.AuxInt != 63 || !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(smagic(64, c).m) && s == smagic(64, c).s && x.Op != OpConst64 && sdivisibleOK(64, c)) {
							continue
						}
						v.reset(OpLeq64U)
						v0 := b.NewValue0(v.Pos, OpRotateLeft64, typ.UInt64)
						v1 := b.NewValue0(v.Pos, OpAdd64, typ.UInt64)
						v2 := b.NewValue0(v.Pos, OpMul64, typ.UInt64)
						v3 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
						v3.AuxInt = int64(sdivisible(64, c).m)
						v2.AddArg2(v3, x)
						v4 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
						v4.AuxInt = int64(sdivisible(64, c).a)
						v1.AddArg2(v2, v4)
						v5 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
						v5.AuxInt = int64(64 - sdivisible(64, c).k)
						v0.AddArg2(v1, v5)
						v6 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
						v6.AuxInt = int64(sdivisible(64, c).max)
						v.AddArg2(v0, v6)
						return true
					}
				}
			}
		}
		break
	}
	// match: (Eq64 n (Lsh64x64 (Rsh64x64 (Add64 <t> n (Rsh64Ux64 <t> (Rsh64x64 <t> n (Const64 <typ.UInt64> [63])) (Const64 <typ.UInt64> [kbar]))) (Const64 <typ.UInt64> [k])) (Const64 <typ.UInt64> [k])) )
	// cond: k > 0 && k < 63 && kbar == 64 - k
	// result: (Eq64 (And64 <t> n (Const64 <t> [int64(1<<uint(k)-1)])) (Const64 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpLsh64x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpRsh64x64 {
				continue
			}
			_ = v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAdd64 {
				continue
			}
			t := v_1_0_0.Type
			_ = v_1_0_0.Args[1]
			v_1_0_0_0 := v_1_0_0.Args[0]
			v_1_0_0_1 := v_1_0_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0_0_0, v_1_0_0_1 = _i1+1, v_1_0_0_1, v_1_0_0_0 {
				if n != v_1_0_0_0 || v_1_0_0_1.Op != OpRsh64Ux64 || v_1_0_0_1.Type != t {
					continue
				}
				_ = v_1_0_0_1.Args[1]
				v_1_0_0_1_0 := v_1_0_0_1.Args[0]
				if v_1_0_0_1_0.Op != OpRsh64x64 || v_1_0_0_1_0.Type != t {
					continue
				}
				_ = v_1_0_0_1_0.Args[1]
				if n != v_1_0_0_1_0.Args[0] {
					continue
				}
				v_1_0_0_1_0_1 := v_1_0_0_1_0.Args[1]
				if v_1_0_0_1_0_1.Op != OpConst64 || v_1_0_0_1_0_1.Type != typ.UInt64 || v_1_0_0_1_0_1.AuxInt != 63 {
					continue
				}
				v_1_0_0_1_1 := v_1_0_0_1.Args[1]
				if v_1_0_0_1_1.Op != OpConst64 || v_1_0_0_1_1.Type != typ.UInt64 {
					continue
				}
				kbar := v_1_0_0_1_1.AuxInt
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpConst64 || v_1_0_1.Type != typ.UInt64 {
					continue
				}
				k := v_1_0_1.AuxInt
				v_1_1 := v_1.Args[1]
				if v_1_1.Op != OpConst64 || v_1_1.Type != typ.UInt64 || v_1_1.AuxInt != k || !(k > 0 && k < 63 && kbar == 64-k) {
					continue
				}
				v.reset(OpEq64)
				v0 := b.NewValue0(v.Pos, OpAnd64, t)
				v1 := b.NewValue0(v.Pos, OpConst64, t)
				v1.AuxInt = int64(1<<uint(k) - 1)
				v0.AddArg2(n, v1)
				v2 := b.NewValue0(v.Pos, OpConst64, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	// match: (Eq64 s:(Sub64 x y) (Const64 [0]))
	// cond: s.Uses == 1
	// result: (Eq64 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s := v_0
			if s.Op != OpSub64 {
				continue
			}
			y := s.Args[1]
			x := s.Args[0]
			if v_1.Op != OpConst64 || v_1.AuxInt != 0 || !(s.Uses == 1) {
				continue
			}
			v.reset(OpEq64)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (Eq64 (And64 <t> x (Const64 <t> [y])) (Const64 <t> [y]))
	// cond: isPowerOfTwo(y)
	// result: (Neq64 (And64 <t> x (Const64 <t> [y])) (Const64 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd64 {
				continue
			}
			t := v_0.Type
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst64 || v_0_1.Type != t {
					continue
				}
				y := v_0_1.AuxInt
				if v_1.Op != OpConst64 || v_1.Type != t || v_1.AuxInt != y || !(isPowerOfTwo(y)) {
					continue
				}
				v.reset(OpNeq64)
				v0 := b.NewValue0(v.Pos, OpAnd64, t)
				v1 := b.NewValue0(v.Pos, OpConst64, t)
				v1.AuxInt = y
				v0.AddArg2(x, v1)
				v2 := b.NewValue0(v.Pos, OpConst64, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpEq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Eq64F (Const64F [c]) (Const64F [d]))
	// result: (ConstBool [b2i(auxTo64F(c) == auxTo64F(d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64F {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64F {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(auxTo64F(c) == auxTo64F(d))
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Eq8 x x)
	// result: (ConstBool [1])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (Eq8 (Const8 <t> [c]) (Add8 (Const8 <t> [d]) x))
	// result: (Eq8 (Const8 <t> [int64(int8(c-d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst8 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpEq8)
				v0 := b.NewValue0(v.Pos, OpConst8, t)
				v0.AuxInt = int64(int8(c - d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Eq8 (Const8 [c]) (Const8 [d]))
	// result: (ConstBool [b2i(c == d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst8 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c == d)
			return true
		}
		break
	}
	// match: (Eq8 (Mod8u x (Const8 [c])) (Const8 [0]))
	// cond: x.Op != OpConst8 && udivisibleOK(8,c) && !hasSmallRotate(config)
	// result: (Eq32 (Mod32u <typ.UInt32> (ZeroExt8to32 <typ.UInt32> x) (Const32 <typ.UInt32> [c&0xff])) (Const32 <typ.UInt32> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMod8u {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpConst8 {
				continue
			}
			c := v_0_1.AuxInt
			if v_1.Op != OpConst8 || v_1.AuxInt != 0 || !(x.Op != OpConst8 && udivisibleOK(8, c) && !hasSmallRotate(config)) {
				continue
			}
			v.reset(OpEq32)
			v0 := b.NewValue0(v.Pos, OpMod32u, typ.UInt32)
			v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
			v1.AddArg(x)
			v2 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
			v2.AuxInt = c & 0xff
			v0.AddArg2(v1, v2)
			v3 := b.NewValue0(v.Pos, OpConst32, typ.UInt32)
			v3.AuxInt = 0
			v.AddArg2(v0, v3)
			return true
		}
		break
	}
	// match: (Eq8 (Mod8 x (Const8 [c])) (Const8 [0]))
	// cond: x.Op != OpConst8 && sdivisibleOK(8,c) && !hasSmallRotate(config)
	// result: (Eq32 (Mod32 <typ.Int32> (SignExt8to32 <typ.Int32> x) (Const32 <typ.Int32> [c])) (Const32 <typ.Int32> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpMod8 {
				continue
			}
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpConst8 {
				continue
			}
			c := v_0_1.AuxInt
			if v_1.Op != OpConst8 || v_1.AuxInt != 0 || !(x.Op != OpConst8 && sdivisibleOK(8, c) && !hasSmallRotate(config)) {
				continue
			}
			v.reset(OpEq32)
			v0 := b.NewValue0(v.Pos, OpMod32, typ.Int32)
			v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
			v1.AddArg(x)
			v2 := b.NewValue0(v.Pos, OpConst32, typ.Int32)
			v2.AuxInt = c
			v0.AddArg2(v1, v2)
			v3 := b.NewValue0(v.Pos, OpConst32, typ.Int32)
			v3.AuxInt = 0
			v.AddArg2(v0, v3)
			return true
		}
		break
	}
	// match: (Eq8 x (Mul8 (Const8 [c]) (Trunc32to8 (Rsh32Ux64 mul:(Mul32 (Const32 [m]) (ZeroExt8to32 x)) (Const64 [s]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<8+umagic(8,c).m) && s == 8+umagic(8,c).s && x.Op != OpConst8 && udivisibleOK(8,c)
	// result: (Leq8U (RotateLeft8 <typ.UInt8> (Mul8 <typ.UInt8> (Const8 <typ.UInt8> [int64(int8(udivisible(8,c).m))]) x) (Const8 <typ.UInt8> [int64(8-udivisible(8,c).k)]) ) (Const8 <typ.UInt8> [int64(int8(udivisible(8,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst8 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpTrunc32to8 {
					continue
				}
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh32Ux64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				mul := v_1_1_0.Args[0]
				if mul.Op != OpMul32 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst32 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpZeroExt8to32 || x != mul_1.Args[0] {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					if !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(1<<8+umagic(8, c).m) && s == 8+umagic(8, c).s && x.Op != OpConst8 && udivisibleOK(8, c)) {
						continue
					}
					v.reset(OpLeq8U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft8, typ.UInt8)
					v1 := b.NewValue0(v.Pos, OpMul8, typ.UInt8)
					v2 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
					v2.AuxInt = int64(int8(udivisible(8, c).m))
					v1.AddArg2(v2, x)
					v3 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
					v3.AuxInt = int64(8 - udivisible(8, c).k)
					v0.AddArg2(v1, v3)
					v4 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
					v4.AuxInt = int64(int8(udivisible(8, c).max))
					v.AddArg2(v0, v4)
					return true
				}
			}
		}
		break
	}
	// match: (Eq8 x (Mul8 (Const8 [c]) (Sub8 (Rsh32x64 mul:(Mul32 (Const32 [m]) (SignExt8to32 x)) (Const64 [s])) (Rsh32x64 (SignExt8to32 x) (Const64 [31]))) ) )
	// cond: v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(smagic(8,c).m) && s == 8+smagic(8,c).s && x.Op != OpConst8 && sdivisibleOK(8,c)
	// result: (Leq8U (RotateLeft8 <typ.UInt8> (Add8 <typ.UInt8> (Mul8 <typ.UInt8> (Const8 <typ.UInt8> [int64(int8(sdivisible(8,c).m))]) x) (Const8 <typ.UInt8> [int64(int8(sdivisible(8,c).a))]) ) (Const8 <typ.UInt8> [int64(8-sdivisible(8,c).k)]) ) (Const8 <typ.UInt8> [int64(int8(sdivisible(8,c).max))]) )
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMul8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst8 {
					continue
				}
				c := v_1_0.AuxInt
				if v_1_1.Op != OpSub8 {
					continue
				}
				_ = v_1_1.Args[1]
				v_1_1_0 := v_1_1.Args[0]
				if v_1_1_0.Op != OpRsh32x64 {
					continue
				}
				_ = v_1_1_0.Args[1]
				mul := v_1_1_0.Args[0]
				if mul.Op != OpMul32 {
					continue
				}
				_ = mul.Args[1]
				mul_0 := mul.Args[0]
				mul_1 := mul.Args[1]
				for _i2 := 0; _i2 <= 1; _i2, mul_0, mul_1 = _i2+1, mul_1, mul_0 {
					if mul_0.Op != OpConst32 {
						continue
					}
					m := mul_0.AuxInt
					if mul_1.Op != OpSignExt8to32 || x != mul_1.Args[0] {
						continue
					}
					v_1_1_0_1 := v_1_1_0.Args[1]
					if v_1_1_0_1.Op != OpConst64 {
						continue
					}
					s := v_1_1_0_1.AuxInt
					v_1_1_1 := v_1_1.Args[1]
					if v_1_1_1.Op != OpRsh32x64 {
						continue
					}
					_ = v_1_1_1.Args[1]
					v_1_1_1_0 := v_1_1_1.Args[0]
					if v_1_1_1_0.Op != OpSignExt8to32 || x != v_1_1_1_0.Args[0] {
						continue
					}
					v_1_1_1_1 := v_1_1_1.Args[1]
					if v_1_1_1_1.Op != OpConst64 || v_1_1_1_1.AuxInt != 31 || !(v.Block.Func.pass.name != "opt" && mul.Uses == 1 && m == int64(smagic(8, c).m) && s == 8+smagic(8, c).s && x.Op != OpConst8 && sdivisibleOK(8, c)) {
						continue
					}
					v.reset(OpLeq8U)
					v0 := b.NewValue0(v.Pos, OpRotateLeft8, typ.UInt8)
					v1 := b.NewValue0(v.Pos, OpAdd8, typ.UInt8)
					v2 := b.NewValue0(v.Pos, OpMul8, typ.UInt8)
					v3 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
					v3.AuxInt = int64(int8(sdivisible(8, c).m))
					v2.AddArg2(v3, x)
					v4 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
					v4.AuxInt = int64(int8(sdivisible(8, c).a))
					v1.AddArg2(v2, v4)
					v5 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
					v5.AuxInt = int64(8 - sdivisible(8, c).k)
					v0.AddArg2(v1, v5)
					v6 := b.NewValue0(v.Pos, OpConst8, typ.UInt8)
					v6.AuxInt = int64(int8(sdivisible(8, c).max))
					v.AddArg2(v0, v6)
					return true
				}
			}
		}
		break
	}
	// match: (Eq8 n (Lsh8x64 (Rsh8x64 (Add8 <t> n (Rsh8Ux64 <t> (Rsh8x64 <t> n (Const64 <typ.UInt64> [ 7])) (Const64 <typ.UInt64> [kbar]))) (Const64 <typ.UInt64> [k])) (Const64 <typ.UInt64> [k])) )
	// cond: k > 0 && k < 7 && kbar == 8 - k
	// result: (Eq8 (And8 <t> n (Const8 <t> [int64(1<<uint(k)-1)])) (Const8 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpLsh8x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpRsh8x64 {
				continue
			}
			_ = v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAdd8 {
				continue
			}
			t := v_1_0_0.Type
			_ = v_1_0_0.Args[1]
			v_1_0_0_0 := v_1_0_0.Args[0]
			v_1_0_0_1 := v_1_0_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0_0_0, v_1_0_0_1 = _i1+1, v_1_0_0_1, v_1_0_0_0 {
				if n != v_1_0_0_0 || v_1_0_0_1.Op != OpRsh8Ux64 || v_1_0_0_1.Type != t {
					continue
				}
				_ = v_1_0_0_1.Args[1]
				v_1_0_0_1_0 := v_1_0_0_1.Args[0]
				if v_1_0_0_1_0.Op != OpRsh8x64 || v_1_0_0_1_0.Type != t {
					continue
				}
				_ = v_1_0_0_1_0.Args[1]
				if n != v_1_0_0_1_0.Args[0] {
					continue
				}
				v_1_0_0_1_0_1 := v_1_0_0_1_0.Args[1]
				if v_1_0_0_1_0_1.Op != OpConst64 || v_1_0_0_1_0_1.Type != typ.UInt64 || v_1_0_0_1_0_1.AuxInt != 7 {
					continue
				}
				v_1_0_0_1_1 := v_1_0_0_1.Args[1]
				if v_1_0_0_1_1.Op != OpConst64 || v_1_0_0_1_1.Type != typ.UInt64 {
					continue
				}
				kbar := v_1_0_0_1_1.AuxInt
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpConst64 || v_1_0_1.Type != typ.UInt64 {
					continue
				}
				k := v_1_0_1.AuxInt
				v_1_1 := v_1.Args[1]
				if v_1_1.Op != OpConst64 || v_1_1.Type != typ.UInt64 || v_1_1.AuxInt != k || !(k > 0 && k < 7 && kbar == 8-k) {
					continue
				}
				v.reset(OpEq8)
				v0 := b.NewValue0(v.Pos, OpAnd8, t)
				v1 := b.NewValue0(v.Pos, OpConst8, t)
				v1.AuxInt = int64(1<<uint(k) - 1)
				v0.AddArg2(n, v1)
				v2 := b.NewValue0(v.Pos, OpConst8, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	// match: (Eq8 s:(Sub8 x y) (Const8 [0]))
	// cond: s.Uses == 1
	// result: (Eq8 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s := v_0
			if s.Op != OpSub8 {
				continue
			}
			y := s.Args[1]
			x := s.Args[0]
			if v_1.Op != OpConst8 || v_1.AuxInt != 0 || !(s.Uses == 1) {
				continue
			}
			v.reset(OpEq8)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (Eq8 (And8 <t> x (Const8 <t> [y])) (Const8 <t> [y]))
	// cond: isPowerOfTwo(y)
	// result: (Neq8 (And8 <t> x (Const8 <t> [y])) (Const8 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd8 {
				continue
			}
			t := v_0.Type
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst8 || v_0_1.Type != t {
					continue
				}
				y := v_0_1.AuxInt
				if v_1.Op != OpConst8 || v_1.Type != t || v_1.AuxInt != y || !(isPowerOfTwo(y)) {
					continue
				}
				v.reset(OpNeq8)
				v0 := b.NewValue0(v.Pos, OpAnd8, t)
				v1 := b.NewValue0(v.Pos, OpConst8, t)
				v1.AuxInt = y
				v0.AddArg2(x, v1)
				v2 := b.NewValue0(v.Pos, OpConst8, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpEqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (EqB (ConstBool [c]) (ConstBool [d]))
	// result: (ConstBool [b2i(c == d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConstBool {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConstBool {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c == d)
			return true
		}
		break
	}
	// match: (EqB (ConstBool [0]) x)
	// result: (Not x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConstBool || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.reset(OpNot)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (EqB (ConstBool [1]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConstBool || v_0.AuxInt != 1 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpEqInter(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqInter x y)
	// result: (EqPtr (ITab x) (ITab y))
	for {
		x := v_0
		y := v_1
		v.reset(OpEqPtr)
		v0 := b.NewValue0(v.Pos, OpITab, typ.Uintptr)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpITab, typ.Uintptr)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValuegeneric_OpEqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqPtr x x)
	// result: (ConstBool [1])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (EqPtr (Addr {a} _) (Addr {b} _))
	// result: (ConstBool [b2i(a == b)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAddr {
				continue
			}
			a := v_0.Aux
			if v_1.Op != OpAddr {
				continue
			}
			b := v_1.Aux
			v.reset(OpConstBool)
			v.AuxInt = b2i(a == b)
			return true
		}
		break
	}
	// match: (EqPtr (Addr {a} _) (OffPtr [o] (Addr {b} _)))
	// result: (ConstBool [b2i(a == b && o == 0)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAddr {
				continue
			}
			a := v_0.Aux
			if v_1.Op != OpOffPtr {
				continue
			}
			o := v_1.AuxInt
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpAddr {
				continue
			}
			b := v_1_0.Aux
			v.reset(OpConstBool)
			v.AuxInt = b2i(a == b && o == 0)
			return true
		}
		break
	}
	// match: (EqPtr (OffPtr [o1] (Addr {a} _)) (OffPtr [o2] (Addr {b} _)))
	// result: (ConstBool [b2i(a == b && o1 == o2)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			o1 := v_0.AuxInt
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAddr {
				continue
			}
			a := v_0_0.Aux
			if v_1.Op != OpOffPtr {
				continue
			}
			o2 := v_1.AuxInt
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpAddr {
				continue
			}
			b := v_1_0.Aux
			v.reset(OpConstBool)
			v.AuxInt = b2i(a == b && o1 == o2)
			return true
		}
		break
	}
	// match: (EqPtr (LocalAddr {a} _ _) (LocalAddr {b} _ _))
	// result: (ConstBool [b2i(a == b)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLocalAddr {
				continue
			}
			a := v_0.Aux
			_ = v_0.Args[1]
			if v_1.Op != OpLocalAddr {
				continue
			}
			b := v_1.Aux
			_ = v_1.Args[1]
			v.reset(OpConstBool)
			v.AuxInt = b2i(a == b)
			return true
		}
		break
	}
	// match: (EqPtr (LocalAddr {a} _ _) (OffPtr [o] (LocalAddr {b} _ _)))
	// result: (ConstBool [b2i(a == b && o == 0)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLocalAddr {
				continue
			}
			a := v_0.Aux
			_ = v_0.Args[1]
			if v_1.Op != OpOffPtr {
				continue
			}
			o := v_1.AuxInt
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpLocalAddr {
				continue
			}
			b := v_1_0.Aux
			_ = v_1_0.Args[1]
			v.reset(OpConstBool)
			v.AuxInt = b2i(a == b && o == 0)
			return true
		}
		break
	}
	// match: (EqPtr (OffPtr [o1] (LocalAddr {a} _ _)) (OffPtr [o2] (LocalAddr {b} _ _)))
	// result: (ConstBool [b2i(a == b && o1 == o2)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			o1 := v_0.AuxInt
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpLocalAddr {
				continue
			}
			a := v_0_0.Aux
			_ = v_0_0.Args[1]
			if v_1.Op != OpOffPtr {
				continue
			}
			o2 := v_1.AuxInt
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpLocalAddr {
				continue
			}
			b := v_1_0.Aux
			_ = v_1_0.Args[1]
			v.reset(OpConstBool)
			v.AuxInt = b2i(a == b && o1 == o2)
			return true
		}
		break
	}
	// match: (EqPtr (OffPtr [o1] p1) p2)
	// cond: isSamePtr(p1, p2)
	// result: (ConstBool [b2i(o1 == 0)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			o1 := v_0.AuxInt
			p1 := v_0.Args[0]
			p2 := v_1
			if !(isSamePtr(p1, p2)) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = b2i(o1 == 0)
			return true
		}
		break
	}
	// match: (EqPtr (OffPtr [o1] p1) (OffPtr [o2] p2))
	// cond: isSamePtr(p1, p2)
	// result: (ConstBool [b2i(o1 == o2)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			o1 := v_0.AuxInt
			p1 := v_0.Args[0]
			if v_1.Op != OpOffPtr {
				continue
			}
			o2 := v_1.AuxInt
			p2 := v_1.Args[0]
			if !(isSamePtr(p1, p2)) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = b2i(o1 == o2)
			return true
		}
		break
	}
	// match: (EqPtr (Const32 [c]) (Const32 [d]))
	// result: (ConstBool [b2i(c == d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c == d)
			return true
		}
		break
	}
	// match: (EqPtr (Const64 [c]) (Const64 [d]))
	// result: (ConstBool [b2i(c == d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c == d)
			return true
		}
		break
	}
	// match: (EqPtr (LocalAddr _ _) (Addr _))
	// result: (ConstBool [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLocalAddr {
				continue
			}
			_ = v_0.Args[1]
			if v_1.Op != OpAddr {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (EqPtr (OffPtr (LocalAddr _ _)) (Addr _))
	// result: (ConstBool [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpLocalAddr {
				continue
			}
			_ = v_0_0.Args[1]
			if v_1.Op != OpAddr {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (EqPtr (LocalAddr _ _) (OffPtr (Addr _)))
	// result: (ConstBool [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLocalAddr {
				continue
			}
			_ = v_0.Args[1]
			if v_1.Op != OpOffPtr {
				continue
			}
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpAddr {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (EqPtr (OffPtr (LocalAddr _ _)) (OffPtr (Addr _)))
	// result: (ConstBool [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpLocalAddr {
				continue
			}
			_ = v_0_0.Args[1]
			if v_1.Op != OpOffPtr {
				continue
			}
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpAddr {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (EqPtr (AddPtr p1 o1) p2)
	// cond: isSamePtr(p1, p2)
	// result: (Not (IsNonNil o1))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAddPtr {
				continue
			}
			o1 := v_0.Args[1]
			p1 := v_0.Args[0]
			p2 := v_1
			if !(isSamePtr(p1, p2)) {
				continue
			}
			v.reset(OpNot)
			v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
			v0.AddArg(o1)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (EqPtr (Const32 [0]) p)
	// result: (Not (IsNonNil p))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
				continue
			}
			p := v_1
			v.reset(OpNot)
			v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
			v0.AddArg(p)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (EqPtr (Const64 [0]) p)
	// result: (Not (IsNonNil p))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
				continue
			}
			p := v_1
			v.reset(OpNot)
			v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
			v0.AddArg(p)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (EqPtr (ConstNil) p)
	// result: (Not (IsNonNil p))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConstNil {
				continue
			}
			p := v_1
			v.reset(OpNot)
			v0 := b.NewValue0(v.Pos, OpIsNonNil, typ.Bool)
			v0.AddArg(p)
			v.AddArg(v0)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpEqSlice(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqSlice x y)
	// result: (EqPtr (SlicePtr x) (SlicePtr y))
	for {
		x := v_0
		y := v_1
		v.reset(OpEqPtr)
		v0 := b.NewValue0(v.Pos, OpSlicePtr, typ.BytePtr)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSlicePtr, typ.BytePtr)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValuegeneric_OpGeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Geq32F (Const32F [c]) (Const32F [d]))
	// result: (ConstBool [b2i(auxTo32F(c) >= auxTo32F(d))])
	for {
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(auxTo32F(c) >= auxTo32F(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Geq64F (Const64F [c]) (Const64F [d]))
	// result: (ConstBool [b2i(auxTo64F(c) >= auxTo64F(d))])
	for {
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(auxTo64F(c) >= auxTo64F(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Greater32F (Const32F [c]) (Const32F [d]))
	// result: (ConstBool [b2i(auxTo32F(c) > auxTo32F(d))])
	for {
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(auxTo32F(c) > auxTo32F(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpGreater64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Greater64F (Const64F [c]) (Const64F [d]))
	// result: (ConstBool [b2i(auxTo64F(c) > auxTo64F(d))])
	for {
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(auxTo64F(c) > auxTo64F(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpIMake(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (IMake typ (StructMake1 val))
	// result: (IMake typ val)
	for {
		typ := v_0
		if v_1.Op != OpStructMake1 {
			break
		}
		val := v_1.Args[0]
		v.reset(OpIMake)
		v.AddArg2(typ, val)
		return true
	}
	// match: (IMake typ (ArrayMake1 val))
	// result: (IMake typ val)
	for {
		typ := v_0
		if v_1.Op != OpArrayMake1 {
			break
		}
		val := v_1.Args[0]
		v.reset(OpIMake)
		v.AddArg2(typ, val)
		return true
	}
	return false
}
func rewriteValuegeneric_OpInterCall(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (InterCall [argsize] (Load (OffPtr [off] (ITab (IMake (Addr {itab} (SB)) _))) _) mem)
	// cond: devirt(v, itab, off) != nil
	// result: (StaticCall [argsize] {devirt(v, itab, off)} mem)
	for {
		argsize := v.AuxInt
		if v_0.Op != OpLoad {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpOffPtr {
			break
		}
		off := v_0_0.AuxInt
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpITab {
			break
		}
		v_0_0_0_0 := v_0_0_0.Args[0]
		if v_0_0_0_0.Op != OpIMake {
			break
		}
		_ = v_0_0_0_0.Args[1]
		v_0_0_0_0_0 := v_0_0_0_0.Args[0]
		if v_0_0_0_0_0.Op != OpAddr {
			break
		}
		itab := v_0_0_0_0_0.Aux
		v_0_0_0_0_0_0 := v_0_0_0_0_0.Args[0]
		if v_0_0_0_0_0_0.Op != OpSB {
			break
		}
		mem := v_1
		if !(devirt(v, itab, off) != nil) {
			break
		}
		v.reset(OpStaticCall)
		v.AuxInt = argsize
		v.Aux = devirt(v, itab, off)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValuegeneric_OpIsInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (IsInBounds (ZeroExt8to32 _) (Const32 [c]))
	// cond: (1 << 8) <= c
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt8to32 || v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !((1 << 8) <= c) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (ZeroExt8to64 _) (Const64 [c]))
	// cond: (1 << 8) <= c
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt8to64 || v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !((1 << 8) <= c) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (ZeroExt16to32 _) (Const32 [c]))
	// cond: (1 << 16) <= c
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt16to32 || v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !((1 << 16) <= c) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (ZeroExt16to64 _) (Const64 [c]))
	// cond: (1 << 16) <= c
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt16to64 || v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !((1 << 16) <= c) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds x x)
	// result: (ConstBool [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}
	// match: (IsInBounds (And8 (Const8 [c]) _) (Const8 [d]))
	// cond: 0 <= c && c < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpAnd8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst8 {
				continue
			}
			c := v_0_0.AuxInt
			if v_1.Op != OpConst8 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c < d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsInBounds (ZeroExt8to16 (And8 (Const8 [c]) _)) (Const16 [d]))
	// cond: 0 <= c && c < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt8to16 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd8 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0_0, v_0_0_1 = _i0+1, v_0_0_1, v_0_0_0 {
			if v_0_0_0.Op != OpConst8 {
				continue
			}
			c := v_0_0_0.AuxInt
			if v_1.Op != OpConst16 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c < d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsInBounds (ZeroExt8to32 (And8 (Const8 [c]) _)) (Const32 [d]))
	// cond: 0 <= c && c < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt8to32 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd8 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0_0, v_0_0_1 = _i0+1, v_0_0_1, v_0_0_0 {
			if v_0_0_0.Op != OpConst8 {
				continue
			}
			c := v_0_0_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c < d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsInBounds (ZeroExt8to64 (And8 (Const8 [c]) _)) (Const64 [d]))
	// cond: 0 <= c && c < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd8 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0_0, v_0_0_1 = _i0+1, v_0_0_1, v_0_0_0 {
			if v_0_0_0.Op != OpConst8 {
				continue
			}
			c := v_0_0_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c < d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsInBounds (And16 (Const16 [c]) _) (Const16 [d]))
	// cond: 0 <= c && c < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst16 {
				continue
			}
			c := v_0_0.AuxInt
			if v_1.Op != OpConst16 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c < d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsInBounds (ZeroExt16to32 (And16 (Const16 [c]) _)) (Const32 [d]))
	// cond: 0 <= c && c < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt16to32 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd16 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0_0, v_0_0_1 = _i0+1, v_0_0_1, v_0_0_0 {
			if v_0_0_0.Op != OpConst16 {
				continue
			}
			c := v_0_0_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c < d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsInBounds (ZeroExt16to64 (And16 (Const16 [c]) _)) (Const64 [d]))
	// cond: 0 <= c && c < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt16to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd16 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0_0, v_0_0_1 = _i0+1, v_0_0_1, v_0_0_0 {
			if v_0_0_0.Op != OpConst16 {
				continue
			}
			c := v_0_0_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c < d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsInBounds (And32 (Const32 [c]) _) (Const32 [d]))
	// cond: 0 <= c && c < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst32 {
				continue
			}
			c := v_0_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c < d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsInBounds (ZeroExt32to64 (And32 (Const32 [c]) _)) (Const64 [d]))
	// cond: 0 <= c && c < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt32to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAnd32 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0_0, v_0_0_1 = _i0+1, v_0_0_1, v_0_0_0 {
			if v_0_0_0.Op != OpConst32 {
				continue
			}
			c := v_0_0_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c < d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsInBounds (And64 (Const64 [c]) _) (Const64 [d]))
	// cond: 0 <= c && c < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst64 {
				continue
			}
			c := v_0_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c < d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsInBounds (Const32 [c]) (Const32 [d]))
	// result: (ConstBool [b2i(0 <= c && c < d)])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(0 <= c && c < d)
		return true
	}
	// match: (IsInBounds (Const64 [c]) (Const64 [d]))
	// result: (ConstBool [b2i(0 <= c && c < d)])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(0 <= c && c < d)
		return true
	}
	// match: (IsInBounds (Mod32u _ y) y)
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpMod32u {
			break
		}
		y := v_0.Args[1]
		if y != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (Mod64u _ y) y)
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpMod64u {
			break
		}
		y := v_0.Args[1]
		if y != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (ZeroExt8to64 (Rsh8Ux64 _ (Const64 [c]))) (Const64 [d]))
	// cond: 0 < c && c < 8 && 1<<uint( 8-c)-1 < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 8 && 1<<uint(8-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (ZeroExt8to32 (Rsh8Ux64 _ (Const64 [c]))) (Const32 [d]))
	// cond: 0 < c && c < 8 && 1<<uint( 8-c)-1 < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt8to32 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 8 && 1<<uint(8-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (ZeroExt8to16 (Rsh8Ux64 _ (Const64 [c]))) (Const16 [d]))
	// cond: 0 < c && c < 8 && 1<<uint( 8-c)-1 < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt8to16 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 8 && 1<<uint(8-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (Rsh8Ux64 _ (Const64 [c])) (Const64 [d]))
	// cond: 0 < c && c < 8 && 1<<uint( 8-c)-1 < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 8 && 1<<uint(8-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (ZeroExt16to64 (Rsh16Ux64 _ (Const64 [c]))) (Const64 [d]))
	// cond: 0 < c && c < 16 && 1<<uint(16-c)-1 < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt16to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 16 && 1<<uint(16-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (ZeroExt16to32 (Rsh16Ux64 _ (Const64 [c]))) (Const64 [d]))
	// cond: 0 < c && c < 16 && 1<<uint(16-c)-1 < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt16to32 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 16 && 1<<uint(16-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (Rsh16Ux64 _ (Const64 [c])) (Const64 [d]))
	// cond: 0 < c && c < 16 && 1<<uint(16-c)-1 < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 16 && 1<<uint(16-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (ZeroExt32to64 (Rsh32Ux64 _ (Const64 [c]))) (Const64 [d]))
	// cond: 0 < c && c < 32 && 1<<uint(32-c)-1 < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpZeroExt32to64 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh32Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c := v_0_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 32 && 1<<uint(32-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (Rsh32Ux64 _ (Const64 [c])) (Const64 [d]))
	// cond: 0 < c && c < 32 && 1<<uint(32-c)-1 < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpRsh32Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 32 && 1<<uint(32-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsInBounds (Rsh64Ux64 _ (Const64 [c])) (Const64 [d]))
	// cond: 0 < c && c < 64 && 1<<uint(64-c)-1 < d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpRsh64Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(0 < c && c < 64 && 1<<uint(64-c)-1 < d) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	// match: (IsNonNil (ConstNil))
	// result: (ConstBool [0])
	for {
		if v_0.Op != OpConstNil {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}
	// match: (IsNonNil (Const32 [c]))
	// result: (ConstBool [b2i(c != 0)])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != 0)
		return true
	}
	// match: (IsNonNil (Const64 [c]))
	// result: (ConstBool [b2i(c != 0)])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c != 0)
		return true
	}
	// match: (IsNonNil (Addr _))
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpAddr {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsNonNil (LocalAddr _ _))
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpLocalAddr {
			break
		}
		_ = v_0.Args[1]
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpIsSliceInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (IsSliceInBounds x x)
	// result: (ConstBool [1])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsSliceInBounds (And32 (Const32 [c]) _) (Const32 [d]))
	// cond: 0 <= c && c <= d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst32 {
				continue
			}
			c := v_0_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c <= d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsSliceInBounds (And64 (Const64 [c]) _) (Const64 [d]))
	// cond: 0 <= c && c <= d
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst64 {
				continue
			}
			c := v_0_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			if !(0 <= c && c <= d) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (IsSliceInBounds (Const32 [0]) _)
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsSliceInBounds (Const64 [0]) _)
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	// match: (IsSliceInBounds (Const32 [c]) (Const32 [d]))
	// result: (ConstBool [b2i(0 <= c && c <= d)])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(0 <= c && c <= d)
		return true
	}
	// match: (IsSliceInBounds (Const64 [c]) (Const64 [d]))
	// result: (ConstBool [b2i(0 <= c && c <= d)])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(0 <= c && c <= d)
		return true
	}
	// match: (IsSliceInBounds (SliceLen x) (SliceCap x))
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpSliceLen {
			break
		}
		x := v_0.Args[0]
		if v_1.Op != OpSliceCap || x != v_1.Args[0] {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq16 (Const16 [c]) (Const16 [d]))
	// result: (ConstBool [b2i(c <= d)])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c <= d)
		return true
	}
	// match: (Leq16 (Const16 [0]) (And16 _ (Const16 [c])))
	// cond: int16(c) >= 0
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 || v_1.Op != OpAnd16 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			if v_1_1.Op != OpConst16 {
				continue
			}
			c := v_1_1.AuxInt
			if !(int16(c) >= 0) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (Leq16 (Const16 [0]) (Rsh16Ux64 _ (Const64 [c])))
	// cond: c > 0
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 || v_1.Op != OpRsh16Ux64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c > 0) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq16U (Const16 [c]) (Const16 [d]))
	// result: (ConstBool [b2i(uint16(c) <= uint16(d))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint16(c) <= uint16(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq32 (Const32 [c]) (Const32 [d]))
	// result: (ConstBool [b2i(c <= d)])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c <= d)
		return true
	}
	// match: (Leq32 (Const32 [0]) (And32 _ (Const32 [c])))
	// cond: int32(c) >= 0
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 || v_1.Op != OpAnd32 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			if v_1_1.Op != OpConst32 {
				continue
			}
			c := v_1_1.AuxInt
			if !(int32(c) >= 0) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (Leq32 (Const32 [0]) (Rsh32Ux64 _ (Const64 [c])))
	// cond: c > 0
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 || v_1.Op != OpRsh32Ux64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c > 0) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq32F (Const32F [c]) (Const32F [d]))
	// result: (ConstBool [b2i(auxTo32F(c) <= auxTo32F(d))])
	for {
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(auxTo32F(c) <= auxTo32F(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq32U (Const32 [c]) (Const32 [d]))
	// result: (ConstBool [b2i(uint32(c) <= uint32(d))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint32(c) <= uint32(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq64 (Const64 [c]) (Const64 [d]))
	// result: (ConstBool [b2i(c <= d)])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c <= d)
		return true
	}
	// match: (Leq64 (Const64 [0]) (And64 _ (Const64 [c])))
	// cond: int64(c) >= 0
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 || v_1.Op != OpAnd64 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			if v_1_1.Op != OpConst64 {
				continue
			}
			c := v_1_1.AuxInt
			if !(int64(c) >= 0) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (Leq64 (Const64 [0]) (Rsh64Ux64 _ (Const64 [c])))
	// cond: c > 0
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 || v_1.Op != OpRsh64Ux64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c > 0) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq64F (Const64F [c]) (Const64F [d]))
	// result: (ConstBool [b2i(auxTo64F(c) <= auxTo64F(d))])
	for {
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(auxTo64F(c) <= auxTo64F(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq64U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq64U (Const64 [c]) (Const64 [d]))
	// result: (ConstBool [b2i(uint64(c) <= uint64(d))])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint64(c) <= uint64(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq8 (Const8 [c]) (Const8 [d]))
	// result: (ConstBool [b2i(c <= d)])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c <= d)
		return true
	}
	// match: (Leq8 (Const8 [0]) (And8 _ (Const8 [c])))
	// cond: int8(c) >= 0
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 || v_1.Op != OpAnd8 {
			break
		}
		_ = v_1.Args[1]
		v_1_0 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_1_0, v_1_1 = _i0+1, v_1_1, v_1_0 {
			if v_1_1.Op != OpConst8 {
				continue
			}
			c := v_1_1.AuxInt
			if !(int8(c) >= 0) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (Leq8 (Const8 [0]) (Rsh8Ux64 _ (Const64 [c])))
	// cond: c > 0
	// result: (ConstBool [1])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 || v_1.Op != OpRsh8Ux64 {
			break
		}
		_ = v_1.Args[1]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		c := v_1_1.AuxInt
		if !(c > 0) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValuegeneric_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Leq8U (Const8 [c]) (Const8 [d]))
	// result: (ConstBool [b2i(uint8(c) <= uint8(d))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint8(c) <= uint8(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less16 (Const16 [c]) (Const16 [d]))
	// result: (ConstBool [b2i(c < d)])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c < d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less16U (Const16 [c]) (Const16 [d]))
	// result: (ConstBool [b2i(uint16(c) < uint16(d))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint16(c) < uint16(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less32 (Const32 [c]) (Const32 [d]))
	// result: (ConstBool [b2i(c < d)])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c < d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less32F (Const32F [c]) (Const32F [d]))
	// result: (ConstBool [b2i(auxTo32F(c) < auxTo32F(d))])
	for {
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(auxTo32F(c) < auxTo32F(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less32U (Const32 [c]) (Const32 [d]))
	// result: (ConstBool [b2i(uint32(c) < uint32(d))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint32(c) < uint32(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less64 (Const64 [c]) (Const64 [d]))
	// result: (ConstBool [b2i(c < d)])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c < d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less64F (Const64F [c]) (Const64F [d]))
	// result: (ConstBool [b2i(auxTo64F(c) < auxTo64F(d))])
	for {
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(auxTo64F(c) < auxTo64F(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess64U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less64U (Const64 [c]) (Const64 [d]))
	// result: (ConstBool [b2i(uint64(c) < uint64(d))])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint64(c) < uint64(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less8 (Const8 [c]) (Const8 [d]))
	// result: (ConstBool [b2i(c < d)])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(c < d)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less8U (Const8 [c]) (Const8 [d]))
	// result: (ConstBool [b2i(uint8(c) < uint8(d))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = b2i(uint8(c) < uint8(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	fe := b.Func.fe
	// match: (Load <t1> p1 (Store {t2} p2 x _))
	// cond: isSamePtr(p1, p2) && t1.Compare(x.Type) == types.CMPeq && t1.Size() == sizeof(t2)
	// result: x
	for {
		t1 := v.Type
		p1 := v_0
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(isSamePtr(p1, p2) && t1.Compare(x.Type) == types.CMPeq && t1.Size() == sizeof(t2)) {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Load <t1> p1 (Store {t2} p2 _ (Store {t3} p3 x _)))
	// cond: isSamePtr(p1, p3) && t1.Compare(x.Type) == types.CMPeq && t1.Size() == sizeof(t2) && disjoint(p3, sizeof(t3), p2, sizeof(t2))
	// result: x
	for {
		t1 := v.Type
		p1 := v_0
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		x := v_1_2.Args[1]
		if !(isSamePtr(p1, p3) && t1.Compare(x.Type) == types.CMPeq && t1.Size() == sizeof(t2) && disjoint(p3, sizeof(t3), p2, sizeof(t2))) {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Load <t1> p1 (Store {t2} p2 _ (Store {t3} p3 _ (Store {t4} p4 x _))))
	// cond: isSamePtr(p1, p4) && t1.Compare(x.Type) == types.CMPeq && t1.Size() == sizeof(t2) && disjoint(p4, sizeof(t4), p2, sizeof(t2)) && disjoint(p4, sizeof(t4), p3, sizeof(t3))
	// result: x
	for {
		t1 := v.Type
		p1 := v_0
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		v_1_2_2 := v_1_2.Args[2]
		if v_1_2_2.Op != OpStore {
			break
		}
		t4 := v_1_2_2.Aux
		_ = v_1_2_2.Args[2]
		p4 := v_1_2_2.Args[0]
		x := v_1_2_2.Args[1]
		if !(isSamePtr(p1, p4) && t1.Compare(x.Type) == types.CMPeq && t1.Size() == sizeof(t2) && disjoint(p4, sizeof(t4), p2, sizeof(t2)) && disjoint(p4, sizeof(t4), p3, sizeof(t3))) {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Load <t1> p1 (Store {t2} p2 _ (Store {t3} p3 _ (Store {t4} p4 _ (Store {t5} p5 x _)))))
	// cond: isSamePtr(p1, p5) && t1.Compare(x.Type) == types.CMPeq && t1.Size() == sizeof(t2) && disjoint(p5, sizeof(t5), p2, sizeof(t2)) && disjoint(p5, sizeof(t5), p3, sizeof(t3)) && disjoint(p5, sizeof(t5), p4, sizeof(t4))
	// result: x
	for {
		t1 := v.Type
		p1 := v_0
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		v_1_2_2 := v_1_2.Args[2]
		if v_1_2_2.Op != OpStore {
			break
		}
		t4 := v_1_2_2.Aux
		_ = v_1_2_2.Args[2]
		p4 := v_1_2_2.Args[0]
		v_1_2_2_2 := v_1_2_2.Args[2]
		if v_1_2_2_2.Op != OpStore {
			break
		}
		t5 := v_1_2_2_2.Aux
		_ = v_1_2_2_2.Args[2]
		p5 := v_1_2_2_2.Args[0]
		x := v_1_2_2_2.Args[1]
		if !(isSamePtr(p1, p5) && t1.Compare(x.Type) == types.CMPeq && t1.Size() == sizeof(t2) && disjoint(p5, sizeof(t5), p2, sizeof(t2)) && disjoint(p5, sizeof(t5), p3, sizeof(t3)) && disjoint(p5, sizeof(t5), p4, sizeof(t4))) {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Load <t1> p1 (Store {t2} p2 (Const64 [x]) _))
	// cond: isSamePtr(p1,p2) && sizeof(t2) == 8 && is64BitFloat(t1)
	// result: (Const64F [x])
	for {
		t1 := v.Type
		p1 := v_0
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 {
			break
		}
		x := v_1_1.AuxInt
		if !(isSamePtr(p1, p2) && sizeof(t2) == 8 && is64BitFloat(t1)) {
			break
		}
		v.reset(OpConst64F)
		v.AuxInt = x
		return true
	}
	// match: (Load <t1> p1 (Store {t2} p2 (Const32 [x]) _))
	// cond: isSamePtr(p1,p2) && sizeof(t2) == 4 && is32BitFloat(t1)
	// result: (Const32F [auxFrom32F(math.Float32frombits(uint32(x)))])
	for {
		t1 := v.Type
		p1 := v_0
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 {
			break
		}
		x := v_1_1.AuxInt
		if !(isSamePtr(p1, p2) && sizeof(t2) == 4 && is32BitFloat(t1)) {
			break
		}
		v.reset(OpConst32F)
		v.AuxInt = auxFrom32F(math.Float32frombits(uint32(x)))
		return true
	}
	// match: (Load <t1> p1 (Store {t2} p2 (Const64F [x]) _))
	// cond: isSamePtr(p1,p2) && sizeof(t2) == 8 && is64BitInt(t1)
	// result: (Const64 [x])
	for {
		t1 := v.Type
		p1 := v_0
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64F {
			break
		}
		x := v_1_1.AuxInt
		if !(isSamePtr(p1, p2) && sizeof(t2) == 8 && is64BitInt(t1)) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = x
		return true
	}
	// match: (Load <t1> p1 (Store {t2} p2 (Const32F [x]) _))
	// cond: isSamePtr(p1,p2) && sizeof(t2) == 4 && is32BitInt(t1)
	// result: (Const32 [int64(int32(math.Float32bits(auxTo32F(x))))])
	for {
		t1 := v.Type
		p1 := v_0
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32F {
			break
		}
		x := v_1_1.AuxInt
		if !(isSamePtr(p1, p2) && sizeof(t2) == 4 && is32BitInt(t1)) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int64(int32(math.Float32bits(auxTo32F(x))))
		return true
	}
	// match: (Load <t1> op:(OffPtr [o1] p1) (Store {t2} p2 _ mem:(Zero [n] p3 _)))
	// cond: o1 >= 0 && o1+t1.Size() <= n && isSamePtr(p1, p3) && fe.CanSSA(t1) && disjoint(op, t1.Size(), p2, sizeof(t2))
	// result: @mem.Block (Load <t1> (OffPtr <op.Type> [o1] p3) mem)
	for {
		t1 := v.Type
		op := v_0
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		mem := v_1.Args[2]
		if mem.Op != OpZero {
			break
		}
		n := mem.AuxInt
		_ = mem.Args[1]
		p3 := mem.Args[0]
		if !(o1 >= 0 && o1+t1.Size() <= n && isSamePtr(p1, p3) && fe.CanSSA(t1) && disjoint(op, t1.Size(), p2, sizeof(t2))) {
			break
		}
		b = mem.Block
		v0 := b.NewValue0(v.Pos, OpLoad, t1)
		v.copyOf(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, op.Type)
		v1.AuxInt = o1
		v1.AddArg(p3)
		v0.AddArg2(v1, mem)
		return true
	}
	// match: (Load <t1> op:(OffPtr [o1] p1) (Store {t2} p2 _ (Store {t3} p3 _ mem:(Zero [n] p4 _))))
	// cond: o1 >= 0 && o1+t1.Size() <= n && isSamePtr(p1, p4) && fe.CanSSA(t1) && disjoint(op, t1.Size(), p2, sizeof(t2)) && disjoint(op, t1.Size(), p3, sizeof(t3))
	// result: @mem.Block (Load <t1> (OffPtr <op.Type> [o1] p4) mem)
	for {
		t1 := v.Type
		op := v_0
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		mem := v_1_2.Args[2]
		if mem.Op != OpZero {
			break
		}
		n := mem.AuxInt
		_ = mem.Args[1]
		p4 := mem.Args[0]
		if !(o1 >= 0 && o1+t1.Size() <= n && isSamePtr(p1, p4) && fe.CanSSA(t1) && disjoint(op, t1.Size(), p2, sizeof(t2)) && disjoint(op, t1.Size(), p3, sizeof(t3))) {
			break
		}
		b = mem.Block
		v0 := b.NewValue0(v.Pos, OpLoad, t1)
		v.copyOf(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, op.Type)
		v1.AuxInt = o1
		v1.AddArg(p4)
		v0.AddArg2(v1, mem)
		return true
	}
	// match: (Load <t1> op:(OffPtr [o1] p1) (Store {t2} p2 _ (Store {t3} p3 _ (Store {t4} p4 _ mem:(Zero [n] p5 _)))))
	// cond: o1 >= 0 && o1+t1.Size() <= n && isSamePtr(p1, p5) && fe.CanSSA(t1) && disjoint(op, t1.Size(), p2, sizeof(t2)) && disjoint(op, t1.Size(), p3, sizeof(t3)) && disjoint(op, t1.Size(), p4, sizeof(t4))
	// result: @mem.Block (Load <t1> (OffPtr <op.Type> [o1] p5) mem)
	for {
		t1 := v.Type
		op := v_0
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		v_1_2_2 := v_1_2.Args[2]
		if v_1_2_2.Op != OpStore {
			break
		}
		t4 := v_1_2_2.Aux
		_ = v_1_2_2.Args[2]
		p4 := v_1_2_2.Args[0]
		mem := v_1_2_2.Args[2]
		if mem.Op != OpZero {
			break
		}
		n := mem.AuxInt
		_ = mem.Args[1]
		p5 := mem.Args[0]
		if !(o1 >= 0 && o1+t1.Size() <= n && isSamePtr(p1, p5) && fe.CanSSA(t1) && disjoint(op, t1.Size(), p2, sizeof(t2)) && disjoint(op, t1.Size(), p3, sizeof(t3)) && disjoint(op, t1.Size(), p4, sizeof(t4))) {
			break
		}
		b = mem.Block
		v0 := b.NewValue0(v.Pos, OpLoad, t1)
		v.copyOf(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, op.Type)
		v1.AuxInt = o1
		v1.AddArg(p5)
		v0.AddArg2(v1, mem)
		return true
	}
	// match: (Load <t1> op:(OffPtr [o1] p1) (Store {t2} p2 _ (Store {t3} p3 _ (Store {t4} p4 _ (Store {t5} p5 _ mem:(Zero [n] p6 _))))))
	// cond: o1 >= 0 && o1+t1.Size() <= n && isSamePtr(p1, p6) && fe.CanSSA(t1) && disjoint(op, t1.Size(), p2, sizeof(t2)) && disjoint(op, t1.Size(), p3, sizeof(t3)) && disjoint(op, t1.Size(), p4, sizeof(t4)) && disjoint(op, t1.Size(), p5, sizeof(t5))
	// result: @mem.Block (Load <t1> (OffPtr <op.Type> [o1] p6) mem)
	for {
		t1 := v.Type
		op := v_0
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		if v_1.Op != OpStore {
			break
		}
		t2 := v_1.Aux
		_ = v_1.Args[2]
		p2 := v_1.Args[0]
		v_1_2 := v_1.Args[2]
		if v_1_2.Op != OpStore {
			break
		}
		t3 := v_1_2.Aux
		_ = v_1_2.Args[2]
		p3 := v_1_2.Args[0]
		v_1_2_2 := v_1_2.Args[2]
		if v_1_2_2.Op != OpStore {
			break
		}
		t4 := v_1_2_2.Aux
		_ = v_1_2_2.Args[2]
		p4 := v_1_2_2.Args[0]
		v_1_2_2_2 := v_1_2_2.Args[2]
		if v_1_2_2_2.Op != OpStore {
			break
		}
		t5 := v_1_2_2_2.Aux
		_ = v_1_2_2_2.Args[2]
		p5 := v_1_2_2_2.Args[0]
		mem := v_1_2_2_2.Args[2]
		if mem.Op != OpZero {
			break
		}
		n := mem.AuxInt
		_ = mem.Args[1]
		p6 := mem.Args[0]
		if !(o1 >= 0 && o1+t1.Size() <= n && isSamePtr(p1, p6) && fe.CanSSA(t1) && disjoint(op, t1.Size(), p2, sizeof(t2)) && disjoint(op, t1.Size(), p3, sizeof(t3)) && disjoint(op, t1.Size(), p4, sizeof(t4)) && disjoint(op, t1.Size(), p5, sizeof(t5))) {
			break
		}
		b = mem.Block
		v0 := b.NewValue0(v.Pos, OpLoad, t1)
		v.copyOf(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, op.Type)
		v1.AuxInt = o1
		v1.AddArg(p6)
		v0.AddArg2(v1, mem)
		return true
	}
	// match: (Load <t1> (OffPtr [o] p1) (Zero [n] p2 _))
	// cond: t1.IsBoolean() && isSamePtr(p1, p2) && n >= o + 1
	// result: (ConstBool [0])
	for {
		t1 := v.Type
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(t1.IsBoolean() && isSamePtr(p1, p2) && n >= o+1) {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}
	// match: (Load <t1> (OffPtr [o] p1) (Zero [n] p2 _))
	// cond: is8BitInt(t1) && isSamePtr(p1, p2) && n >= o + 1
	// result: (Const8 [0])
	for {
		t1 := v.Type
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(is8BitInt(t1) && isSamePtr(p1, p2) && n >= o+1) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	// match: (Load <t1> (OffPtr [o] p1) (Zero [n] p2 _))
	// cond: is16BitInt(t1) && isSamePtr(p1, p2) && n >= o + 2
	// result: (Const16 [0])
	for {
		t1 := v.Type
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(is16BitInt(t1) && isSamePtr(p1, p2) && n >= o+2) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	// match: (Load <t1> (OffPtr [o] p1) (Zero [n] p2 _))
	// cond: is32BitInt(t1) && isSamePtr(p1, p2) && n >= o + 4
	// result: (Const32 [0])
	for {
		t1 := v.Type
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(is32BitInt(t1) && isSamePtr(p1, p2) && n >= o+4) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	// match: (Load <t1> (OffPtr [o] p1) (Zero [n] p2 _))
	// cond: is64BitInt(t1) && isSamePtr(p1, p2) && n >= o + 8
	// result: (Const64 [0])
	for {
		t1 := v.Type
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(is64BitInt(t1) && isSamePtr(p1, p2) && n >= o+8) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	// match: (Load <t1> (OffPtr [o] p1) (Zero [n] p2 _))
	// cond: is32BitFloat(t1) && isSamePtr(p1, p2) && n >= o + 4
	// result: (Const32F [0])
	for {
		t1 := v.Type
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(is32BitFloat(t1) && isSamePtr(p1, p2) && n >= o+4) {
			break
		}
		v.reset(OpConst32F)
		v.AuxInt = 0
		return true
	}
	// match: (Load <t1> (OffPtr [o] p1) (Zero [n] p2 _))
	// cond: is64BitFloat(t1) && isSamePtr(p1, p2) && n >= o + 8
	// result: (Const64F [0])
	for {
		t1 := v.Type
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		if v_1.Op != OpZero {
			break
		}
		n := v_1.AuxInt
		_ = v_1.Args[1]
		p2 := v_1.Args[0]
		if !(is64BitFloat(t1) && isSamePtr(p1, p2) && n >= o+8) {
			break
		}
		v.reset(OpConst64F)
		v.AuxInt = 0
		return true
	}
	// match: (Load <t> _ _)
	// cond: t.IsStruct() && t.NumFields() == 0 && fe.CanSSA(t)
	// result: (StructMake0)
	for {
		t := v.Type
		if !(t.IsStruct() && t.NumFields() == 0 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake0)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsStruct() && t.NumFields() == 1 && fe.CanSSA(t)
	// result: (StructMake1 (Load <t.FieldType(0)> (OffPtr <t.FieldType(0).PtrTo()> [0] ptr) mem))
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsStruct() && t.NumFields() == 1 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake1)
		v0 := b.NewValue0(v.Pos, OpLoad, t.FieldType(0))
		v1 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(0).PtrTo())
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v0.AddArg2(v1, mem)
		v.AddArg(v0)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsStruct() && t.NumFields() == 2 && fe.CanSSA(t)
	// result: (StructMake2 (Load <t.FieldType(0)> (OffPtr <t.FieldType(0).PtrTo()> [0] ptr) mem) (Load <t.FieldType(1)> (OffPtr <t.FieldType(1).PtrTo()> [t.FieldOff(1)] ptr) mem))
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsStruct() && t.NumFields() == 2 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake2)
		v0 := b.NewValue0(v.Pos, OpLoad, t.FieldType(0))
		v1 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(0).PtrTo())
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v0.AddArg2(v1, mem)
		v2 := b.NewValue0(v.Pos, OpLoad, t.FieldType(1))
		v3 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(1).PtrTo())
		v3.AuxInt = t.FieldOff(1)
		v3.AddArg(ptr)
		v2.AddArg2(v3, mem)
		v.AddArg2(v0, v2)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsStruct() && t.NumFields() == 3 && fe.CanSSA(t)
	// result: (StructMake3 (Load <t.FieldType(0)> (OffPtr <t.FieldType(0).PtrTo()> [0] ptr) mem) (Load <t.FieldType(1)> (OffPtr <t.FieldType(1).PtrTo()> [t.FieldOff(1)] ptr) mem) (Load <t.FieldType(2)> (OffPtr <t.FieldType(2).PtrTo()> [t.FieldOff(2)] ptr) mem))
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsStruct() && t.NumFields() == 3 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake3)
		v0 := b.NewValue0(v.Pos, OpLoad, t.FieldType(0))
		v1 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(0).PtrTo())
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v0.AddArg2(v1, mem)
		v2 := b.NewValue0(v.Pos, OpLoad, t.FieldType(1))
		v3 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(1).PtrTo())
		v3.AuxInt = t.FieldOff(1)
		v3.AddArg(ptr)
		v2.AddArg2(v3, mem)
		v4 := b.NewValue0(v.Pos, OpLoad, t.FieldType(2))
		v5 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(2).PtrTo())
		v5.AuxInt = t.FieldOff(2)
		v5.AddArg(ptr)
		v4.AddArg2(v5, mem)
		v.AddArg3(v0, v2, v4)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsStruct() && t.NumFields() == 4 && fe.CanSSA(t)
	// result: (StructMake4 (Load <t.FieldType(0)> (OffPtr <t.FieldType(0).PtrTo()> [0] ptr) mem) (Load <t.FieldType(1)> (OffPtr <t.FieldType(1).PtrTo()> [t.FieldOff(1)] ptr) mem) (Load <t.FieldType(2)> (OffPtr <t.FieldType(2).PtrTo()> [t.FieldOff(2)] ptr) mem) (Load <t.FieldType(3)> (OffPtr <t.FieldType(3).PtrTo()> [t.FieldOff(3)] ptr) mem))
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsStruct() && t.NumFields() == 4 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpStructMake4)
		v0 := b.NewValue0(v.Pos, OpLoad, t.FieldType(0))
		v1 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(0).PtrTo())
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v0.AddArg2(v1, mem)
		v2 := b.NewValue0(v.Pos, OpLoad, t.FieldType(1))
		v3 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(1).PtrTo())
		v3.AuxInt = t.FieldOff(1)
		v3.AddArg(ptr)
		v2.AddArg2(v3, mem)
		v4 := b.NewValue0(v.Pos, OpLoad, t.FieldType(2))
		v5 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(2).PtrTo())
		v5.AuxInt = t.FieldOff(2)
		v5.AddArg(ptr)
		v4.AddArg2(v5, mem)
		v6 := b.NewValue0(v.Pos, OpLoad, t.FieldType(3))
		v7 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(3).PtrTo())
		v7.AuxInt = t.FieldOff(3)
		v7.AddArg(ptr)
		v6.AddArg2(v7, mem)
		v.AddArg4(v0, v2, v4, v6)
		return true
	}
	// match: (Load <t> _ _)
	// cond: t.IsArray() && t.NumElem() == 0
	// result: (ArrayMake0)
	for {
		t := v.Type
		if !(t.IsArray() && t.NumElem() == 0) {
			break
		}
		v.reset(OpArrayMake0)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsArray() && t.NumElem() == 1 && fe.CanSSA(t)
	// result: (ArrayMake1 (Load <t.Elem()> ptr mem))
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsArray() && t.NumElem() == 1 && fe.CanSSA(t)) {
			break
		}
		v.reset(OpArrayMake1)
		v0 := b.NewValue0(v.Pos, OpLoad, t.Elem())
		v0.AddArg2(ptr, mem)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh16x16 <t> x (Const16 [c]))
	// result: (Lsh16x64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh16x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh16x16 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh16x32 <t> x (Const32 [c]))
	// result: (Lsh16x64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh16x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh16x32 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x64 (Const16 [c]) (Const64 [d]))
	// result: (Const16 [int64(int16(c) << uint64(d))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c) << uint64(d))
		return true
	}
	// match: (Lsh16x64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Lsh16x64 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh16x64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (Const16 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh16x64 <t> (Lsh16x64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Lsh16x64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpLsh16x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpLsh16x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh16x64 (Rsh16Ux64 (Lsh16x64 x (Const64 [c1])) (Const64 [c2])) (Const64 [c3]))
	// cond: uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)
	// result: (Lsh16x64 x (Const64 <typ.UInt64> [c1-c2+c3]))
	for {
		if v_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLsh16x64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpLsh16x64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh16x8 <t> x (Const8 [c]))
	// result: (Lsh16x64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh16x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh16x8 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh32x16 <t> x (Const16 [c]))
	// result: (Lsh32x64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh32x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh32x16 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh32x32 <t> x (Const32 [c]))
	// result: (Lsh32x64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh32x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh32x32 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x64 (Const32 [c]) (Const64 [d]))
	// result: (Const32 [int64(int32(c) << uint64(d))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c) << uint64(d))
		return true
	}
	// match: (Lsh32x64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Lsh32x64 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh32x64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (Const32 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh32x64 <t> (Lsh32x64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Lsh32x64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpLsh32x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh32x64 (Rsh32Ux64 (Lsh32x64 x (Const64 [c1])) (Const64 [c2])) (Const64 [c3]))
	// cond: uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)
	// result: (Lsh32x64 x (Const64 <typ.UInt64> [c1-c2+c3]))
	for {
		if v_0.Op != OpRsh32Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpLsh32x64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh32x8 <t> x (Const8 [c]))
	// result: (Lsh32x64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh32x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh32x8 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh64x16 <t> x (Const16 [c]))
	// result: (Lsh64x64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh64x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh64x16 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh64x32 <t> x (Const32 [c]))
	// result: (Lsh64x64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh64x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh64x32 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x64 (Const64 [c]) (Const64 [d]))
	// result: (Const64 [c << uint64(d)])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c << uint64(d)
		return true
	}
	// match: (Lsh64x64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Lsh64x64 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh64x64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (Const64 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh64x64 <t> (Lsh64x64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Lsh64x64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpLsh64x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh64x64 (Rsh64Ux64 (Lsh64x64 x (Const64 [c1])) (Const64 [c2])) (Const64 [c3]))
	// cond: uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)
	// result: (Lsh64x64 x (Const64 <typ.UInt64> [c1-c2+c3]))
	for {
		if v_0.Op != OpRsh64Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpLsh64x64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh64x8 <t> x (Const8 [c]))
	// result: (Lsh64x64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh64x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh64x8 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh8x16 <t> x (Const16 [c]))
	// result: (Lsh8x64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh8x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh8x16 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh8x32 <t> x (Const32 [c]))
	// result: (Lsh8x64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh8x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh8x32 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x64 (Const8 [c]) (Const64 [d]))
	// result: (Const8 [int64(int8(c) << uint64(d))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c) << uint64(d))
		return true
	}
	// match: (Lsh8x64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Lsh8x64 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh8x64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (Const8 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	// match: (Lsh8x64 <t> (Lsh8x64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Lsh8x64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpLsh8x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpLsh8x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh8x64 (Rsh8Ux64 (Lsh8x64 x (Const64 [c1])) (Const64 [c2])) (Const64 [c3]))
	// cond: uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)
	// result: (Lsh8x64 x (Const64 <typ.UInt64> [c1-c2+c3]))
	for {
		if v_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLsh8x64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpLsh8x64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh8x8 <t> x (Const8 [c]))
	// result: (Lsh8x64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpLsh8x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh8x8 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Mod16 (Const16 [c]) (Const16 [d]))
	// cond: d != 0
	// result: (Const16 [int64(int16(c % d))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c % d))
		return true
	}
	// match: (Mod16 <t> n (Const16 [c]))
	// cond: isNonNegative(n) && isPowerOfTwo(c&0xffff)
	// result: (And16 n (Const16 <t> [(c&0xffff)-1]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xffff)) {
			break
		}
		v.reset(OpAnd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = (c & 0xffff) - 1
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod16 <t> n (Const16 [c]))
	// cond: c < 0 && c != -1<<15
	// result: (Mod16 <t> n (Const16 <t> [-c]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<15) {
			break
		}
		v.reset(OpMod16)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = -c
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod16 <t> x (Const16 [c]))
	// cond: x.Op != OpConst16 && (c > 0 || c == -1<<15)
	// result: (Sub16 x (Mul16 <t> (Div16 <t> x (Const16 <t> [c])) (Const16 <t> [c])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst16 && (c > 0 || c == -1<<15)) {
			break
		}
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpMul16, t)
		v1 := b.NewValue0(v.Pos, OpDiv16, t)
		v2 := b.NewValue0(v.Pos, OpConst16, t)
		v2.AuxInt = c
		v1.AddArg2(x, v2)
		v3 := b.NewValue0(v.Pos, OpConst16, t)
		v3.AuxInt = c
		v0.AddArg2(v1, v3)
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Mod16u (Const16 [c]) (Const16 [d]))
	// cond: d != 0
	// result: (Const16 [int64(uint16(c) % uint16(d))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = int64(uint16(c) % uint16(d))
		return true
	}
	// match: (Mod16u <t> n (Const16 [c]))
	// cond: isPowerOfTwo(c&0xffff)
	// result: (And16 n (Const16 <t> [(c&0xffff)-1]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xffff)) {
			break
		}
		v.reset(OpAnd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = (c & 0xffff) - 1
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod16u <t> x (Const16 [c]))
	// cond: x.Op != OpConst16 && c > 0 && umagicOK(16,c)
	// result: (Sub16 x (Mul16 <t> (Div16u <t> x (Const16 <t> [c])) (Const16 <t> [c])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst16 && c > 0 && umagicOK(16, c)) {
			break
		}
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpMul16, t)
		v1 := b.NewValue0(v.Pos, OpDiv16u, t)
		v2 := b.NewValue0(v.Pos, OpConst16, t)
		v2.AuxInt = c
		v1.AddArg2(x, v2)
		v3 := b.NewValue0(v.Pos, OpConst16, t)
		v3.AuxInt = c
		v0.AddArg2(v1, v3)
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Mod32 (Const32 [c]) (Const32 [d]))
	// cond: d != 0
	// result: (Const32 [int64(int32(c % d))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c % d))
		return true
	}
	// match: (Mod32 <t> n (Const32 [c]))
	// cond: isNonNegative(n) && isPowerOfTwo(c&0xffffffff)
	// result: (And32 n (Const32 <t> [(c&0xffffffff)-1]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xffffffff)) {
			break
		}
		v.reset(OpAnd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = (c & 0xffffffff) - 1
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod32 <t> n (Const32 [c]))
	// cond: c < 0 && c != -1<<31
	// result: (Mod32 <t> n (Const32 <t> [-c]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<31) {
			break
		}
		v.reset(OpMod32)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = -c
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod32 <t> x (Const32 [c]))
	// cond: x.Op != OpConst32 && (c > 0 || c == -1<<31)
	// result: (Sub32 x (Mul32 <t> (Div32 <t> x (Const32 <t> [c])) (Const32 <t> [c])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst32 && (c > 0 || c == -1<<31)) {
			break
		}
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpMul32, t)
		v1 := b.NewValue0(v.Pos, OpDiv32, t)
		v2 := b.NewValue0(v.Pos, OpConst32, t)
		v2.AuxInt = c
		v1.AddArg2(x, v2)
		v3 := b.NewValue0(v.Pos, OpConst32, t)
		v3.AuxInt = c
		v0.AddArg2(v1, v3)
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Mod32u (Const32 [c]) (Const32 [d]))
	// cond: d != 0
	// result: (Const32 [int64(uint32(c) % uint32(d))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = int64(uint32(c) % uint32(d))
		return true
	}
	// match: (Mod32u <t> n (Const32 [c]))
	// cond: isPowerOfTwo(c&0xffffffff)
	// result: (And32 n (Const32 <t> [(c&0xffffffff)-1]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xffffffff)) {
			break
		}
		v.reset(OpAnd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = (c & 0xffffffff) - 1
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod32u <t> x (Const32 [c]))
	// cond: x.Op != OpConst32 && c > 0 && umagicOK(32,c)
	// result: (Sub32 x (Mul32 <t> (Div32u <t> x (Const32 <t> [c])) (Const32 <t> [c])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst32 && c > 0 && umagicOK(32, c)) {
			break
		}
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpMul32, t)
		v1 := b.NewValue0(v.Pos, OpDiv32u, t)
		v2 := b.NewValue0(v.Pos, OpConst32, t)
		v2.AuxInt = c
		v1.AddArg2(x, v2)
		v3 := b.NewValue0(v.Pos, OpConst32, t)
		v3.AuxInt = c
		v0.AddArg2(v1, v3)
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Mod64 (Const64 [c]) (Const64 [d]))
	// cond: d != 0
	// result: (Const64 [c % d])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = c % d
		return true
	}
	// match: (Mod64 <t> n (Const64 [c]))
	// cond: isNonNegative(n) && isPowerOfTwo(c)
	// result: (And64 n (Const64 <t> [c-1]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c)) {
			break
		}
		v.reset(OpAnd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - 1
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod64 n (Const64 [-1<<63]))
	// cond: isNonNegative(n)
	// result: n
	for {
		n := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != -1<<63 || !(isNonNegative(n)) {
			break
		}
		v.copyOf(n)
		return true
	}
	// match: (Mod64 <t> n (Const64 [c]))
	// cond: c < 0 && c != -1<<63
	// result: (Mod64 <t> n (Const64 <t> [-c]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<63) {
			break
		}
		v.reset(OpMod64)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = -c
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod64 <t> x (Const64 [c]))
	// cond: x.Op != OpConst64 && (c > 0 || c == -1<<63)
	// result: (Sub64 x (Mul64 <t> (Div64 <t> x (Const64 <t> [c])) (Const64 <t> [c])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst64 && (c > 0 || c == -1<<63)) {
			break
		}
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpMul64, t)
		v1 := b.NewValue0(v.Pos, OpDiv64, t)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = c
		v1.AddArg2(x, v2)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = c
		v0.AddArg2(v1, v3)
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Mod64u (Const64 [c]) (Const64 [d]))
	// cond: d != 0
	// result: (Const64 [int64(uint64(c) % uint64(d))])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = int64(uint64(c) % uint64(d))
		return true
	}
	// match: (Mod64u <t> n (Const64 [c]))
	// cond: isPowerOfTwo(c)
	// result: (And64 n (Const64 <t> [c-1]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpAnd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - 1
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod64u <t> n (Const64 [-1<<63]))
	// result: (And64 n (Const64 <t> [1<<63-1]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != -1<<63 {
			break
		}
		v.reset(OpAnd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = 1<<63 - 1
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod64u <t> x (Const64 [c]))
	// cond: x.Op != OpConst64 && c > 0 && umagicOK(64,c)
	// result: (Sub64 x (Mul64 <t> (Div64u <t> x (Const64 <t> [c])) (Const64 <t> [c])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst64 && c > 0 && umagicOK(64, c)) {
			break
		}
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpMul64, t)
		v1 := b.NewValue0(v.Pos, OpDiv64u, t)
		v2 := b.NewValue0(v.Pos, OpConst64, t)
		v2.AuxInt = c
		v1.AddArg2(x, v2)
		v3 := b.NewValue0(v.Pos, OpConst64, t)
		v3.AuxInt = c
		v0.AddArg2(v1, v3)
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Mod8 (Const8 [c]) (Const8 [d]))
	// cond: d != 0
	// result: (Const8 [int64(int8(c % d))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c % d))
		return true
	}
	// match: (Mod8 <t> n (Const8 [c]))
	// cond: isNonNegative(n) && isPowerOfTwo(c&0xff)
	// result: (And8 n (Const8 <t> [(c&0xff)-1]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(isNonNegative(n) && isPowerOfTwo(c&0xff)) {
			break
		}
		v.reset(OpAnd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = (c & 0xff) - 1
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod8 <t> n (Const8 [c]))
	// cond: c < 0 && c != -1<<7
	// result: (Mod8 <t> n (Const8 <t> [-c]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(c < 0 && c != -1<<7) {
			break
		}
		v.reset(OpMod8)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = -c
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod8 <t> x (Const8 [c]))
	// cond: x.Op != OpConst8 && (c > 0 || c == -1<<7)
	// result: (Sub8 x (Mul8 <t> (Div8 <t> x (Const8 <t> [c])) (Const8 <t> [c])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst8 && (c > 0 || c == -1<<7)) {
			break
		}
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpMul8, t)
		v1 := b.NewValue0(v.Pos, OpDiv8, t)
		v2 := b.NewValue0(v.Pos, OpConst8, t)
		v2.AuxInt = c
		v1.AddArg2(x, v2)
		v3 := b.NewValue0(v.Pos, OpConst8, t)
		v3.AuxInt = c
		v0.AddArg2(v1, v3)
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Mod8u (Const8 [c]) (Const8 [d]))
	// cond: d != 0
	// result: (Const8 [int64(uint8(c) % uint8(d))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		if !(d != 0) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = int64(uint8(c) % uint8(d))
		return true
	}
	// match: (Mod8u <t> n (Const8 [c]))
	// cond: isPowerOfTwo(c&0xff)
	// result: (And8 n (Const8 <t> [(c&0xff)-1]))
	for {
		t := v.Type
		n := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(isPowerOfTwo(c & 0xff)) {
			break
		}
		v.reset(OpAnd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = (c & 0xff) - 1
		v.AddArg2(n, v0)
		return true
	}
	// match: (Mod8u <t> x (Const8 [c]))
	// cond: x.Op != OpConst8 && c > 0 && umagicOK(8, c)
	// result: (Sub8 x (Mul8 <t> (Div8u <t> x (Const8 <t> [c])) (Const8 <t> [c])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(x.Op != OpConst8 && c > 0 && umagicOK(8, c)) {
			break
		}
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpMul8, t)
		v1 := b.NewValue0(v.Pos, OpDiv8u, t)
		v2 := b.NewValue0(v.Pos, OpConst8, t)
		v2.AuxInt = c
		v1.AddArg2(x, v2)
		v3 := b.NewValue0(v.Pos, OpConst8, t)
		v3.AuxInt = c
		v0.AddArg2(v1, v3)
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (Move {t} [n] dst1 src mem:(Zero {t} [n] dst2 _))
	// cond: isSamePtr(src, dst2)
	// result: (Zero {t} [n] dst1 mem)
	for {
		n := v.AuxInt
		t := v.Aux
		dst1 := v_0
		src := v_1
		mem := v_2
		if mem.Op != OpZero || mem.AuxInt != n || mem.Aux != t {
			break
		}
		_ = mem.Args[1]
		dst2 := mem.Args[0]
		if !(isSamePtr(src, dst2)) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = n
		v.Aux = t
		v.AddArg2(dst1, mem)
		return true
	}
	// match: (Move {t} [n] dst1 src mem:(VarDef (Zero {t} [n] dst0 _)))
	// cond: isSamePtr(src, dst0)
	// result: (Zero {t} [n] dst1 mem)
	for {
		n := v.AuxInt
		t := v.Aux
		dst1 := v_0
		src := v_1
		mem := v_2
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpZero || mem_0.AuxInt != n || mem_0.Aux != t {
			break
		}
		_ = mem_0.Args[1]
		dst0 := mem_0.Args[0]
		if !(isSamePtr(src, dst0)) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = n
		v.Aux = t
		v.AddArg2(dst1, mem)
		return true
	}
	// match: (Move {t1} [n] dst1 src1 store:(Store {t2} op:(OffPtr [o2] dst2) _ mem))
	// cond: isSamePtr(dst1, dst2) && store.Uses == 1 && n >= o2 + sizeof(t2) && disjoint(src1, n, op, sizeof(t2)) && clobber(store)
	// result: (Move {t1} [n] dst1 src1 mem)
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst1 := v_0
		src1 := v_1
		store := v_2
		if store.Op != OpStore {
			break
		}
		t2 := store.Aux
		mem := store.Args[2]
		op := store.Args[0]
		if op.Op != OpOffPtr {
			break
		}
		o2 := op.AuxInt
		dst2 := op.Args[0]
		if !(isSamePtr(dst1, dst2) && store.Uses == 1 && n >= o2+sizeof(t2) && disjoint(src1, n, op, sizeof(t2)) && clobber(store)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = n
		v.Aux = t1
		v.AddArg3(dst1, src1, mem)
		return true
	}
	// match: (Move {t} [n] dst1 src1 move:(Move {t} [n] dst2 _ mem))
	// cond: move.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(move)
	// result: (Move {t} [n] dst1 src1 mem)
	for {
		n := v.AuxInt
		t := v.Aux
		dst1 := v_0
		src1 := v_1
		move := v_2
		if move.Op != OpMove || move.AuxInt != n || move.Aux != t {
			break
		}
		mem := move.Args[2]
		dst2 := move.Args[0]
		if !(move.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(move)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = n
		v.Aux = t
		v.AddArg3(dst1, src1, mem)
		return true
	}
	// match: (Move {t} [n] dst1 src1 vardef:(VarDef {x} move:(Move {t} [n] dst2 _ mem)))
	// cond: move.Uses == 1 && vardef.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(move, vardef)
	// result: (Move {t} [n] dst1 src1 (VarDef {x} mem))
	for {
		n := v.AuxInt
		t := v.Aux
		dst1 := v_0
		src1 := v_1
		vardef := v_2
		if vardef.Op != OpVarDef {
			break
		}
		x := vardef.Aux
		move := vardef.Args[0]
		if move.Op != OpMove || move.AuxInt != n || move.Aux != t {
			break
		}
		mem := move.Args[2]
		dst2 := move.Args[0]
		if !(move.Uses == 1 && vardef.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(move, vardef)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = n
		v.Aux = t
		v0 := b.NewValue0(v.Pos, OpVarDef, types.TypeMem)
		v0.Aux = x
		v0.AddArg(mem)
		v.AddArg3(dst1, src1, v0)
		return true
	}
	// match: (Move {t} [n] dst1 src1 zero:(Zero {t} [n] dst2 mem))
	// cond: zero.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(zero)
	// result: (Move {t} [n] dst1 src1 mem)
	for {
		n := v.AuxInt
		t := v.Aux
		dst1 := v_0
		src1 := v_1
		zero := v_2
		if zero.Op != OpZero || zero.AuxInt != n || zero.Aux != t {
			break
		}
		mem := zero.Args[1]
		dst2 := zero.Args[0]
		if !(zero.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(zero)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = n
		v.Aux = t
		v.AddArg3(dst1, src1, mem)
		return true
	}
	// match: (Move {t} [n] dst1 src1 vardef:(VarDef {x} zero:(Zero {t} [n] dst2 mem)))
	// cond: zero.Uses == 1 && vardef.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(zero, vardef)
	// result: (Move {t} [n] dst1 src1 (VarDef {x} mem))
	for {
		n := v.AuxInt
		t := v.Aux
		dst1 := v_0
		src1 := v_1
		vardef := v_2
		if vardef.Op != OpVarDef {
			break
		}
		x := vardef.Aux
		zero := vardef.Args[0]
		if zero.Op != OpZero || zero.AuxInt != n || zero.Aux != t {
			break
		}
		mem := zero.Args[1]
		dst2 := zero.Args[0]
		if !(zero.Uses == 1 && vardef.Uses == 1 && isSamePtr(dst1, dst2) && disjoint(src1, n, dst2, n) && clobber(zero, vardef)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = n
		v.Aux = t
		v0 := b.NewValue0(v.Pos, OpVarDef, types.TypeMem)
		v0.Aux = x
		v0.AddArg(mem)
		v.AddArg3(dst1, src1, v0)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(Store {t2} op2:(OffPtr <tt2> [o2] p2) d1 (Store {t3} op3:(OffPtr <tt3> [0] p3) d2 _)))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && o2 == sizeof(t3) && n == sizeof(t2) + sizeof(t3)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [0] dst) d2 mem))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		op2 := mem.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		tt2 := op2.Type
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		op3 := mem_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		tt3 := op3.Type
		if op3.AuxInt != 0 {
			break
		}
		p3 := op3.Args[0]
		d2 := mem_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && o2 == sizeof(t3) && n == sizeof(t2)+sizeof(t3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = 0
		v2.AddArg(dst)
		v1.AddArg3(v2, d2, mem)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(Store {t2} op2:(OffPtr <tt2> [o2] p2) d1 (Store {t3} op3:(OffPtr <tt3> [o3] p3) d2 (Store {t4} op4:(OffPtr <tt4> [0] p4) d3 _))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && o3 == sizeof(t4) && o2-o3 == sizeof(t3) && n == sizeof(t2) + sizeof(t3) + sizeof(t4)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [o3] dst) d2 (Store {t4} (OffPtr <tt4> [0] dst) d3 mem)))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		op2 := mem.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		tt2 := op2.Type
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		op3 := mem_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		tt3 := op3.Type
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d2 := mem_2.Args[1]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		op4 := mem_2_2.Args[0]
		if op4.Op != OpOffPtr {
			break
		}
		tt4 := op4.Type
		if op4.AuxInt != 0 {
			break
		}
		p4 := op4.Args[0]
		d3 := mem_2_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && o3 == sizeof(t4) && o2-o3 == sizeof(t3) && n == sizeof(t2)+sizeof(t3)+sizeof(t4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = 0
		v4.AddArg(dst)
		v3.AddArg3(v4, d3, mem)
		v1.AddArg3(v2, d2, v3)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(Store {t2} op2:(OffPtr <tt2> [o2] p2) d1 (Store {t3} op3:(OffPtr <tt3> [o3] p3) d2 (Store {t4} op4:(OffPtr <tt4> [o4] p4) d3 (Store {t5} op5:(OffPtr <tt5> [0] p5) d4 _)))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && registerizable(b, t5) && o4 == sizeof(t5) && o3-o4 == sizeof(t4) && o2-o3 == sizeof(t3) && n == sizeof(t2) + sizeof(t3) + sizeof(t4) + sizeof(t5)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [o3] dst) d2 (Store {t4} (OffPtr <tt4> [o4] dst) d3 (Store {t5} (OffPtr <tt5> [0] dst) d4 mem))))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		op2 := mem.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		tt2 := op2.Type
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		op3 := mem_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		tt3 := op3.Type
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d2 := mem_2.Args[1]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		op4 := mem_2_2.Args[0]
		if op4.Op != OpOffPtr {
			break
		}
		tt4 := op4.Type
		o4 := op4.AuxInt
		p4 := op4.Args[0]
		d3 := mem_2_2.Args[1]
		mem_2_2_2 := mem_2_2.Args[2]
		if mem_2_2_2.Op != OpStore {
			break
		}
		t5 := mem_2_2_2.Aux
		_ = mem_2_2_2.Args[2]
		op5 := mem_2_2_2.Args[0]
		if op5.Op != OpOffPtr {
			break
		}
		tt5 := op5.Type
		if op5.AuxInt != 0 {
			break
		}
		p5 := op5.Args[0]
		d4 := mem_2_2_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && registerizable(b, t5) && o4 == sizeof(t5) && o3-o4 == sizeof(t4) && o2-o3 == sizeof(t3) && n == sizeof(t2)+sizeof(t3)+sizeof(t4)+sizeof(t5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = o4
		v4.AddArg(dst)
		v5 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v5.Aux = t5
		v6 := b.NewValue0(v.Pos, OpOffPtr, tt5)
		v6.AuxInt = 0
		v6.AddArg(dst)
		v5.AddArg3(v6, d4, mem)
		v3.AddArg3(v4, d3, v5)
		v1.AddArg3(v2, d2, v3)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(VarDef (Store {t2} op2:(OffPtr <tt2> [o2] p2) d1 (Store {t3} op3:(OffPtr <tt3> [0] p3) d2 _))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && o2 == sizeof(t3) && n == sizeof(t2) + sizeof(t3)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [0] dst) d2 mem))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		op2 := mem_0.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		tt2 := op2.Type
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		op3 := mem_0_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		tt3 := op3.Type
		if op3.AuxInt != 0 {
			break
		}
		p3 := op3.Args[0]
		d2 := mem_0_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && o2 == sizeof(t3) && n == sizeof(t2)+sizeof(t3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = 0
		v2.AddArg(dst)
		v1.AddArg3(v2, d2, mem)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(VarDef (Store {t2} op2:(OffPtr <tt2> [o2] p2) d1 (Store {t3} op3:(OffPtr <tt3> [o3] p3) d2 (Store {t4} op4:(OffPtr <tt4> [0] p4) d3 _)))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && o3 == sizeof(t4) && o2-o3 == sizeof(t3) && n == sizeof(t2) + sizeof(t3) + sizeof(t4)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [o3] dst) d2 (Store {t4} (OffPtr <tt4> [0] dst) d3 mem)))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		op2 := mem_0.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		tt2 := op2.Type
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		op3 := mem_0_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		tt3 := op3.Type
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d2 := mem_0_2.Args[1]
		mem_0_2_2 := mem_0_2.Args[2]
		if mem_0_2_2.Op != OpStore {
			break
		}
		t4 := mem_0_2_2.Aux
		_ = mem_0_2_2.Args[2]
		op4 := mem_0_2_2.Args[0]
		if op4.Op != OpOffPtr {
			break
		}
		tt4 := op4.Type
		if op4.AuxInt != 0 {
			break
		}
		p4 := op4.Args[0]
		d3 := mem_0_2_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && o3 == sizeof(t4) && o2-o3 == sizeof(t3) && n == sizeof(t2)+sizeof(t3)+sizeof(t4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = 0
		v4.AddArg(dst)
		v3.AddArg3(v4, d3, mem)
		v1.AddArg3(v2, d2, v3)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(VarDef (Store {t2} op2:(OffPtr <tt2> [o2] p2) d1 (Store {t3} op3:(OffPtr <tt3> [o3] p3) d2 (Store {t4} op4:(OffPtr <tt4> [o4] p4) d3 (Store {t5} op5:(OffPtr <tt5> [0] p5) d4 _))))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && registerizable(b, t5) && o4 == sizeof(t5) && o3-o4 == sizeof(t4) && o2-o3 == sizeof(t3) && n == sizeof(t2) + sizeof(t3) + sizeof(t4) + sizeof(t5)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [o3] dst) d2 (Store {t4} (OffPtr <tt4> [o4] dst) d3 (Store {t5} (OffPtr <tt5> [0] dst) d4 mem))))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		op2 := mem_0.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		tt2 := op2.Type
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		op3 := mem_0_2.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		tt3 := op3.Type
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d2 := mem_0_2.Args[1]
		mem_0_2_2 := mem_0_2.Args[2]
		if mem_0_2_2.Op != OpStore {
			break
		}
		t4 := mem_0_2_2.Aux
		_ = mem_0_2_2.Args[2]
		op4 := mem_0_2_2.Args[0]
		if op4.Op != OpOffPtr {
			break
		}
		tt4 := op4.Type
		o4 := op4.AuxInt
		p4 := op4.Args[0]
		d3 := mem_0_2_2.Args[1]
		mem_0_2_2_2 := mem_0_2_2.Args[2]
		if mem_0_2_2_2.Op != OpStore {
			break
		}
		t5 := mem_0_2_2_2.Aux
		_ = mem_0_2_2_2.Args[2]
		op5 := mem_0_2_2_2.Args[0]
		if op5.Op != OpOffPtr {
			break
		}
		tt5 := op5.Type
		if op5.AuxInt != 0 {
			break
		}
		p5 := op5.Args[0]
		d4 := mem_0_2_2_2.Args[1]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && registerizable(b, t5) && o4 == sizeof(t5) && o3-o4 == sizeof(t4) && o2-o3 == sizeof(t3) && n == sizeof(t2)+sizeof(t3)+sizeof(t4)+sizeof(t5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = o4
		v4.AddArg(dst)
		v5 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v5.Aux = t5
		v6 := b.NewValue0(v.Pos, OpOffPtr, tt5)
		v6.AuxInt = 0
		v6.AddArg(dst)
		v5.AddArg3(v6, d4, mem)
		v3.AddArg3(v4, d3, v5)
		v1.AddArg3(v2, d2, v3)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(Store {t2} op2:(OffPtr <tt2> [o2] p2) d1 (Zero {t3} [n] p3 _)))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && registerizable(b, t2) && n >= o2 + sizeof(t2)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Zero {t1} [n] dst mem))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		op2 := mem.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		tt2 := op2.Type
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpZero || mem_2.AuxInt != n {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[1]
		p3 := mem_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && registerizable(b, t2) && n >= o2+sizeof(t2)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpZero, types.TypeMem)
		v1.AuxInt = n
		v1.Aux = t1
		v1.AddArg2(dst, mem)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(Store {t2} (OffPtr <tt2> [o2] p2) d1 (Store {t3} (OffPtr <tt3> [o3] p3) d2 (Zero {t4} [n] p4 _))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && n >= o2 + sizeof(t2) && n >= o3 + sizeof(t3)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [o3] dst) d2 (Zero {t1} [n] dst mem)))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		mem_0 := mem.Args[0]
		if mem_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0.Type
		o2 := mem_0.AuxInt
		p2 := mem_0.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		mem_2_0 := mem_2.Args[0]
		if mem_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_2_0.Type
		o3 := mem_2_0.AuxInt
		p3 := mem_2_0.Args[0]
		d2 := mem_2.Args[1]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpZero || mem_2_2.AuxInt != n {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[1]
		p4 := mem_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && n >= o2+sizeof(t2) && n >= o3+sizeof(t3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpZero, types.TypeMem)
		v3.AuxInt = n
		v3.Aux = t1
		v3.AddArg2(dst, mem)
		v1.AddArg3(v2, d2, v3)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(Store {t2} (OffPtr <tt2> [o2] p2) d1 (Store {t3} (OffPtr <tt3> [o3] p3) d2 (Store {t4} (OffPtr <tt4> [o4] p4) d3 (Zero {t5} [n] p5 _)))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && n >= o2 + sizeof(t2) && n >= o3 + sizeof(t3) && n >= o4 + sizeof(t4)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [o3] dst) d2 (Store {t4} (OffPtr <tt4> [o4] dst) d3 (Zero {t1} [n] dst mem))))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		mem_0 := mem.Args[0]
		if mem_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0.Type
		o2 := mem_0.AuxInt
		p2 := mem_0.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		mem_2_0 := mem_2.Args[0]
		if mem_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_2_0.Type
		o3 := mem_2_0.AuxInt
		p3 := mem_2_0.Args[0]
		d2 := mem_2.Args[1]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		mem_2_2_0 := mem_2_2.Args[0]
		if mem_2_2_0.Op != OpOffPtr {
			break
		}
		tt4 := mem_2_2_0.Type
		o4 := mem_2_2_0.AuxInt
		p4 := mem_2_2_0.Args[0]
		d3 := mem_2_2.Args[1]
		mem_2_2_2 := mem_2_2.Args[2]
		if mem_2_2_2.Op != OpZero || mem_2_2_2.AuxInt != n {
			break
		}
		t5 := mem_2_2_2.Aux
		_ = mem_2_2_2.Args[1]
		p5 := mem_2_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && n >= o2+sizeof(t2) && n >= o3+sizeof(t3) && n >= o4+sizeof(t4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = o4
		v4.AddArg(dst)
		v5 := b.NewValue0(v.Pos, OpZero, types.TypeMem)
		v5.AuxInt = n
		v5.Aux = t1
		v5.AddArg2(dst, mem)
		v3.AddArg3(v4, d3, v5)
		v1.AddArg3(v2, d2, v3)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(Store {t2} (OffPtr <tt2> [o2] p2) d1 (Store {t3} (OffPtr <tt3> [o3] p3) d2 (Store {t4} (OffPtr <tt4> [o4] p4) d3 (Store {t5} (OffPtr <tt5> [o5] p5) d4 (Zero {t6} [n] p6 _))))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && isSamePtr(p5, p6) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && alignof(t6) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && registerizable(b, t5) && n >= o2 + sizeof(t2) && n >= o3 + sizeof(t3) && n >= o4 + sizeof(t4) && n >= o5 + sizeof(t5)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [o3] dst) d2 (Store {t4} (OffPtr <tt4> [o4] dst) d3 (Store {t5} (OffPtr <tt5> [o5] dst) d4 (Zero {t1} [n] dst mem)))))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		mem_0 := mem.Args[0]
		if mem_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0.Type
		o2 := mem_0.AuxInt
		p2 := mem_0.Args[0]
		d1 := mem.Args[1]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		mem_2_0 := mem_2.Args[0]
		if mem_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_2_0.Type
		o3 := mem_2_0.AuxInt
		p3 := mem_2_0.Args[0]
		d2 := mem_2.Args[1]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		mem_2_2_0 := mem_2_2.Args[0]
		if mem_2_2_0.Op != OpOffPtr {
			break
		}
		tt4 := mem_2_2_0.Type
		o4 := mem_2_2_0.AuxInt
		p4 := mem_2_2_0.Args[0]
		d3 := mem_2_2.Args[1]
		mem_2_2_2 := mem_2_2.Args[2]
		if mem_2_2_2.Op != OpStore {
			break
		}
		t5 := mem_2_2_2.Aux
		_ = mem_2_2_2.Args[2]
		mem_2_2_2_0 := mem_2_2_2.Args[0]
		if mem_2_2_2_0.Op != OpOffPtr {
			break
		}
		tt5 := mem_2_2_2_0.Type
		o5 := mem_2_2_2_0.AuxInt
		p5 := mem_2_2_2_0.Args[0]
		d4 := mem_2_2_2.Args[1]
		mem_2_2_2_2 := mem_2_2_2.Args[2]
		if mem_2_2_2_2.Op != OpZero || mem_2_2_2_2.AuxInt != n {
			break
		}
		t6 := mem_2_2_2_2.Aux
		_ = mem_2_2_2_2.Args[1]
		p6 := mem_2_2_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && isSamePtr(p5, p6) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && alignof(t6) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && registerizable(b, t5) && n >= o2+sizeof(t2) && n >= o3+sizeof(t3) && n >= o4+sizeof(t4) && n >= o5+sizeof(t5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = o4
		v4.AddArg(dst)
		v5 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v5.Aux = t5
		v6 := b.NewValue0(v.Pos, OpOffPtr, tt5)
		v6.AuxInt = o5
		v6.AddArg(dst)
		v7 := b.NewValue0(v.Pos, OpZero, types.TypeMem)
		v7.AuxInt = n
		v7.Aux = t1
		v7.AddArg2(dst, mem)
		v5.AddArg3(v6, d4, v7)
		v3.AddArg3(v4, d3, v5)
		v1.AddArg3(v2, d2, v3)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(VarDef (Store {t2} op2:(OffPtr <tt2> [o2] p2) d1 (Zero {t3} [n] p3 _))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && registerizable(b, t2) && n >= o2 + sizeof(t2)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Zero {t1} [n] dst mem))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		op2 := mem_0.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		tt2 := op2.Type
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpZero || mem_0_2.AuxInt != n {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[1]
		p3 := mem_0_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && registerizable(b, t2) && n >= o2+sizeof(t2)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpZero, types.TypeMem)
		v1.AuxInt = n
		v1.Aux = t1
		v1.AddArg2(dst, mem)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(VarDef (Store {t2} (OffPtr <tt2> [o2] p2) d1 (Store {t3} (OffPtr <tt3> [o3] p3) d2 (Zero {t4} [n] p4 _)))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && n >= o2 + sizeof(t2) && n >= o3 + sizeof(t3)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [o3] dst) d2 (Zero {t1} [n] dst mem)))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		mem_0_0 := mem_0.Args[0]
		if mem_0_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0_0.Type
		o2 := mem_0_0.AuxInt
		p2 := mem_0_0.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		mem_0_2_0 := mem_0_2.Args[0]
		if mem_0_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_0_2_0.Type
		o3 := mem_0_2_0.AuxInt
		p3 := mem_0_2_0.Args[0]
		d2 := mem_0_2.Args[1]
		mem_0_2_2 := mem_0_2.Args[2]
		if mem_0_2_2.Op != OpZero || mem_0_2_2.AuxInt != n {
			break
		}
		t4 := mem_0_2_2.Aux
		_ = mem_0_2_2.Args[1]
		p4 := mem_0_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && n >= o2+sizeof(t2) && n >= o3+sizeof(t3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpZero, types.TypeMem)
		v3.AuxInt = n
		v3.Aux = t1
		v3.AddArg2(dst, mem)
		v1.AddArg3(v2, d2, v3)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(VarDef (Store {t2} (OffPtr <tt2> [o2] p2) d1 (Store {t3} (OffPtr <tt3> [o3] p3) d2 (Store {t4} (OffPtr <tt4> [o4] p4) d3 (Zero {t5} [n] p5 _))))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && n >= o2 + sizeof(t2) && n >= o3 + sizeof(t3) && n >= o4 + sizeof(t4)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [o3] dst) d2 (Store {t4} (OffPtr <tt4> [o4] dst) d3 (Zero {t1} [n] dst mem))))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		mem_0_0 := mem_0.Args[0]
		if mem_0_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0_0.Type
		o2 := mem_0_0.AuxInt
		p2 := mem_0_0.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		mem_0_2_0 := mem_0_2.Args[0]
		if mem_0_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_0_2_0.Type
		o3 := mem_0_2_0.AuxInt
		p3 := mem_0_2_0.Args[0]
		d2 := mem_0_2.Args[1]
		mem_0_2_2 := mem_0_2.Args[2]
		if mem_0_2_2.Op != OpStore {
			break
		}
		t4 := mem_0_2_2.Aux
		_ = mem_0_2_2.Args[2]
		mem_0_2_2_0 := mem_0_2_2.Args[0]
		if mem_0_2_2_0.Op != OpOffPtr {
			break
		}
		tt4 := mem_0_2_2_0.Type
		o4 := mem_0_2_2_0.AuxInt
		p4 := mem_0_2_2_0.Args[0]
		d3 := mem_0_2_2.Args[1]
		mem_0_2_2_2 := mem_0_2_2.Args[2]
		if mem_0_2_2_2.Op != OpZero || mem_0_2_2_2.AuxInt != n {
			break
		}
		t5 := mem_0_2_2_2.Aux
		_ = mem_0_2_2_2.Args[1]
		p5 := mem_0_2_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && n >= o2+sizeof(t2) && n >= o3+sizeof(t3) && n >= o4+sizeof(t4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = o4
		v4.AddArg(dst)
		v5 := b.NewValue0(v.Pos, OpZero, types.TypeMem)
		v5.AuxInt = n
		v5.Aux = t1
		v5.AddArg2(dst, mem)
		v3.AddArg3(v4, d3, v5)
		v1.AddArg3(v2, d2, v3)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [n] dst p1 mem:(VarDef (Store {t2} (OffPtr <tt2> [o2] p2) d1 (Store {t3} (OffPtr <tt3> [o3] p3) d2 (Store {t4} (OffPtr <tt4> [o4] p4) d3 (Store {t5} (OffPtr <tt5> [o5] p5) d4 (Zero {t6} [n] p6 _)))))))
	// cond: isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && isSamePtr(p5, p6) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && alignof(t6) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && registerizable(b, t5) && n >= o2 + sizeof(t2) && n >= o3 + sizeof(t3) && n >= o4 + sizeof(t4) && n >= o5 + sizeof(t5)
	// result: (Store {t2} (OffPtr <tt2> [o2] dst) d1 (Store {t3} (OffPtr <tt3> [o3] dst) d2 (Store {t4} (OffPtr <tt4> [o4] dst) d3 (Store {t5} (OffPtr <tt5> [o5] dst) d4 (Zero {t1} [n] dst mem)))))
	for {
		n := v.AuxInt
		t1 := v.Aux
		dst := v_0
		p1 := v_1
		mem := v_2
		if mem.Op != OpVarDef {
			break
		}
		mem_0 := mem.Args[0]
		if mem_0.Op != OpStore {
			break
		}
		t2 := mem_0.Aux
		_ = mem_0.Args[2]
		mem_0_0 := mem_0.Args[0]
		if mem_0_0.Op != OpOffPtr {
			break
		}
		tt2 := mem_0_0.Type
		o2 := mem_0_0.AuxInt
		p2 := mem_0_0.Args[0]
		d1 := mem_0.Args[1]
		mem_0_2 := mem_0.Args[2]
		if mem_0_2.Op != OpStore {
			break
		}
		t3 := mem_0_2.Aux
		_ = mem_0_2.Args[2]
		mem_0_2_0 := mem_0_2.Args[0]
		if mem_0_2_0.Op != OpOffPtr {
			break
		}
		tt3 := mem_0_2_0.Type
		o3 := mem_0_2_0.AuxInt
		p3 := mem_0_2_0.Args[0]
		d2 := mem_0_2.Args[1]
		mem_0_2_2 := mem_0_2.Args[2]
		if mem_0_2_2.Op != OpStore {
			break
		}
		t4 := mem_0_2_2.Aux
		_ = mem_0_2_2.Args[2]
		mem_0_2_2_0 := mem_0_2_2.Args[0]
		if mem_0_2_2_0.Op != OpOffPtr {
			break
		}
		tt4 := mem_0_2_2_0.Type
		o4 := mem_0_2_2_0.AuxInt
		p4 := mem_0_2_2_0.Args[0]
		d3 := mem_0_2_2.Args[1]
		mem_0_2_2_2 := mem_0_2_2.Args[2]
		if mem_0_2_2_2.Op != OpStore {
			break
		}
		t5 := mem_0_2_2_2.Aux
		_ = mem_0_2_2_2.Args[2]
		mem_0_2_2_2_0 := mem_0_2_2_2.Args[0]
		if mem_0_2_2_2_0.Op != OpOffPtr {
			break
		}
		tt5 := mem_0_2_2_2_0.Type
		o5 := mem_0_2_2_2_0.AuxInt
		p5 := mem_0_2_2_2_0.Args[0]
		d4 := mem_0_2_2_2.Args[1]
		mem_0_2_2_2_2 := mem_0_2_2_2.Args[2]
		if mem_0_2_2_2_2.Op != OpZero || mem_0_2_2_2_2.AuxInt != n {
			break
		}
		t6 := mem_0_2_2_2_2.Aux
		_ = mem_0_2_2_2_2.Args[1]
		p6 := mem_0_2_2_2_2.Args[0]
		if !(isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && isSamePtr(p5, p6) && alignof(t2) <= alignof(t1) && alignof(t3) <= alignof(t1) && alignof(t4) <= alignof(t1) && alignof(t5) <= alignof(t1) && alignof(t6) <= alignof(t1) && registerizable(b, t2) && registerizable(b, t3) && registerizable(b, t4) && registerizable(b, t5) && n >= o2+sizeof(t2) && n >= o3+sizeof(t3) && n >= o4+sizeof(t4) && n >= o5+sizeof(t5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t2
		v0 := b.NewValue0(v.Pos, OpOffPtr, tt2)
		v0.AuxInt = o2
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpOffPtr, tt3)
		v2.AuxInt = o3
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v3.Aux = t4
		v4 := b.NewValue0(v.Pos, OpOffPtr, tt4)
		v4.AuxInt = o4
		v4.AddArg(dst)
		v5 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v5.Aux = t5
		v6 := b.NewValue0(v.Pos, OpOffPtr, tt5)
		v6.AuxInt = o5
		v6.AddArg(dst)
		v7 := b.NewValue0(v.Pos, OpZero, types.TypeMem)
		v7.AuxInt = n
		v7.Aux = t1
		v7.AddArg2(dst, mem)
		v5.AddArg3(v6, d4, v7)
		v3.AddArg3(v4, d3, v5)
		v1.AddArg3(v2, d2, v3)
		v.AddArg3(v0, d1, v1)
		return true
	}
	// match: (Move {t1} [s] dst tmp1 midmem:(Move {t2} [s] tmp2 src _))
	// cond: t1.(*types.Type).Compare(t2.(*types.Type)) == types.CMPeq && isSamePtr(tmp1, tmp2) && isStackPtr(src) && disjoint(src, s, tmp2, s) && (disjoint(src, s, dst, s) || isInlinableMemmove(dst, src, s, config))
	// result: (Move {t1} [s] dst src midmem)
	for {
		s := v.AuxInt
		t1 := v.Aux
		dst := v_0
		tmp1 := v_1
		midmem := v_2
		if midmem.Op != OpMove || midmem.AuxInt != s {
			break
		}
		t2 := midmem.Aux
		_ = midmem.Args[2]
		tmp2 := midmem.Args[0]
		src := midmem.Args[1]
		if !(t1.(*types.Type).Compare(t2.(*types.Type)) == types.CMPeq && isSamePtr(tmp1, tmp2) && isStackPtr(src) && disjoint(src, s, tmp2, s) && (disjoint(src, s, dst, s) || isInlinableMemmove(dst, src, s, config))) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = s
		v.Aux = t1
		v.AddArg3(dst, src, midmem)
		return true
	}
	// match: (Move {t1} [s] dst tmp1 midmem:(VarDef (Move {t2} [s] tmp2 src _)))
	// cond: t1.(*types.Type).Compare(t2.(*types.Type)) == types.CMPeq && isSamePtr(tmp1, tmp2) && isStackPtr(src) && disjoint(src, s, tmp2, s) && (disjoint(src, s, dst, s) || isInlinableMemmove(dst, src, s, config))
	// result: (Move {t1} [s] dst src midmem)
	for {
		s := v.AuxInt
		t1 := v.Aux
		dst := v_0
		tmp1 := v_1
		midmem := v_2
		if midmem.Op != OpVarDef {
			break
		}
		midmem_0 := midmem.Args[0]
		if midmem_0.Op != OpMove || midmem_0.AuxInt != s {
			break
		}
		t2 := midmem_0.Aux
		_ = midmem_0.Args[2]
		tmp2 := midmem_0.Args[0]
		src := midmem_0.Args[1]
		if !(t1.(*types.Type).Compare(t2.(*types.Type)) == types.CMPeq && isSamePtr(tmp1, tmp2) && isStackPtr(src) && disjoint(src, s, tmp2, s) && (disjoint(src, s, dst, s) || isInlinableMemmove(dst, src, s, config))) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = s
		v.Aux = t1
		v.AddArg3(dst, src, midmem)
		return true
	}
	// match: (Move dst src mem)
	// cond: isSamePtr(dst, src)
	// result: mem
	for {
		dst := v_0
		src := v_1
		mem := v_2
		if !(isSamePtr(dst, src)) {
			break
		}
		v.copyOf(mem)
		return true
	}
	return false
}
func rewriteValuegeneric_OpMul16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mul16 (Const16 [c]) (Const16 [d]))
	// result: (Const16 [int64(int16(c*d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst16 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst16)
			v.AuxInt = int64(int16(c * d))
			return true
		}
		break
	}
	// match: (Mul16 (Const16 [1]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 || v_0.AuxInt != 1 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Mul16 (Const16 [-1]) x)
	// result: (Neg16 x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 || v_0.AuxInt != -1 {
				continue
			}
			x := v_1
			v.reset(OpNeg16)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Mul16 <t> n (Const16 [c]))
	// cond: isPowerOfTwo(c)
	// result: (Lsh16x64 <t> n (Const64 <typ.UInt64> [log2(c)]))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpConst16 {
				continue
			}
			c := v_1.AuxInt
			if !(isPowerOfTwo(c)) {
				continue
			}
			v.reset(OpLsh16x64)
			v.Type = t
			v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
			v0.AuxInt = log2(c)
			v.AddArg2(n, v0)
			return true
		}
		break
	}
	// match: (Mul16 <t> n (Const16 [c]))
	// cond: t.IsSigned() && isPowerOfTwo(-c)
	// result: (Neg16 (Lsh16x64 <t> n (Const64 <typ.UInt64> [log2(-c)])))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpConst16 {
				continue
			}
			c := v_1.AuxInt
			if !(t.IsSigned() && isPowerOfTwo(-c)) {
				continue
			}
			v.reset(OpNeg16)
			v0 := b.NewValue0(v.Pos, OpLsh16x64, t)
			v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
			v1.AuxInt = log2(-c)
			v0.AddArg2(n, v1)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Mul16 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
				continue
			}
			v.reset(OpConst16)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (Mul16 (Const16 <t> [c]) (Mul16 (Const16 <t> [d]) x))
	// result: (Mul16 (Const16 <t> [int64(int16(c*d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpMul16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpMul16)
				v0 := b.NewValue0(v.Pos, OpConst16, t)
				v0.AuxInt = int64(int16(c * d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpMul32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mul32 (Const32 [c]) (Const32 [d]))
	// result: (Const32 [int64(int32(c*d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst32)
			v.AuxInt = int64(int32(c * d))
			return true
		}
		break
	}
	// match: (Mul32 (Const32 [1]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != 1 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Mul32 (Const32 [-1]) x)
	// result: (Neg32 x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != -1 {
				continue
			}
			x := v_1
			v.reset(OpNeg32)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Mul32 <t> n (Const32 [c]))
	// cond: isPowerOfTwo(c)
	// result: (Lsh32x64 <t> n (Const64 <typ.UInt64> [log2(c)]))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpConst32 {
				continue
			}
			c := v_1.AuxInt
			if !(isPowerOfTwo(c)) {
				continue
			}
			v.reset(OpLsh32x64)
			v.Type = t
			v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
			v0.AuxInt = log2(c)
			v.AddArg2(n, v0)
			return true
		}
		break
	}
	// match: (Mul32 <t> n (Const32 [c]))
	// cond: t.IsSigned() && isPowerOfTwo(-c)
	// result: (Neg32 (Lsh32x64 <t> n (Const64 <typ.UInt64> [log2(-c)])))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpConst32 {
				continue
			}
			c := v_1.AuxInt
			if !(t.IsSigned() && isPowerOfTwo(-c)) {
				continue
			}
			v.reset(OpNeg32)
			v0 := b.NewValue0(v.Pos, OpLsh32x64, t)
			v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
			v1.AuxInt = log2(-c)
			v0.AddArg2(n, v1)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Mul32 (Const32 <t> [c]) (Add32 <t> (Const32 <t> [d]) x))
	// result: (Add32 (Const32 <t> [int64(int32(c*d))]) (Mul32 <t> (Const32 <t> [c]) x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd32 || v_1.Type != t {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpAdd32)
				v0 := b.NewValue0(v.Pos, OpConst32, t)
				v0.AuxInt = int64(int32(c * d))
				v1 := b.NewValue0(v.Pos, OpMul32, t)
				v2 := b.NewValue0(v.Pos, OpConst32, t)
				v2.AuxInt = c
				v1.AddArg2(v2, x)
				v.AddArg2(v0, v1)
				return true
			}
		}
		break
	}
	// match: (Mul32 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
				continue
			}
			v.reset(OpConst32)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (Mul32 (Const32 <t> [c]) (Mul32 (Const32 <t> [d]) x))
	// result: (Mul32 (Const32 <t> [int64(int32(c*d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpMul32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpMul32)
				v0 := b.NewValue0(v.Pos, OpConst32, t)
				v0.AuxInt = int64(int32(c * d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpMul32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mul32F (Const32F [c]) (Const32F [d]))
	// result: (Const32F [auxFrom32F(auxTo32F(c) * auxTo32F(d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32F {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32F {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst32F)
			v.AuxInt = auxFrom32F(auxTo32F(c) * auxTo32F(d))
			return true
		}
		break
	}
	// match: (Mul32F x (Const32F [auxFrom64F(1)]))
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpConst32F || v_1.AuxInt != auxFrom64F(1) {
				continue
			}
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Mul32F x (Const32F [auxFrom32F(-1)]))
	// result: (Neg32F x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpConst32F || v_1.AuxInt != auxFrom32F(-1) {
				continue
			}
			v.reset(OpNeg32F)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Mul32F x (Const32F [auxFrom32F(2)]))
	// result: (Add32F x x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpConst32F || v_1.AuxInt != auxFrom32F(2) {
				continue
			}
			v.reset(OpAdd32F)
			v.AddArg2(x, x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpMul64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mul64 (Const64 [c]) (Const64 [d]))
	// result: (Const64 [c*d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst64)
			v.AuxInt = c * d
			return true
		}
		break
	}
	// match: (Mul64 (Const64 [1]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != 1 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Mul64 (Const64 [-1]) x)
	// result: (Neg64 x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != -1 {
				continue
			}
			x := v_1
			v.reset(OpNeg64)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Mul64 <t> n (Const64 [c]))
	// cond: isPowerOfTwo(c)
	// result: (Lsh64x64 <t> n (Const64 <typ.UInt64> [log2(c)]))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpConst64 {
				continue
			}
			c := v_1.AuxInt
			if !(isPowerOfTwo(c)) {
				continue
			}
			v.reset(OpLsh64x64)
			v.Type = t
			v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
			v0.AuxInt = log2(c)
			v.AddArg2(n, v0)
			return true
		}
		break
	}
	// match: (Mul64 <t> n (Const64 [c]))
	// cond: t.IsSigned() && isPowerOfTwo(-c)
	// result: (Neg64 (Lsh64x64 <t> n (Const64 <typ.UInt64> [log2(-c)])))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpConst64 {
				continue
			}
			c := v_1.AuxInt
			if !(t.IsSigned() && isPowerOfTwo(-c)) {
				continue
			}
			v.reset(OpNeg64)
			v0 := b.NewValue0(v.Pos, OpLsh64x64, t)
			v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
			v1.AuxInt = log2(-c)
			v0.AddArg2(n, v1)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Mul64 (Const64 <t> [c]) (Add64 <t> (Const64 <t> [d]) x))
	// result: (Add64 (Const64 <t> [c*d]) (Mul64 <t> (Const64 <t> [c]) x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd64 || v_1.Type != t {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpAdd64)
				v0 := b.NewValue0(v.Pos, OpConst64, t)
				v0.AuxInt = c * d
				v1 := b.NewValue0(v.Pos, OpMul64, t)
				v2 := b.NewValue0(v.Pos, OpConst64, t)
				v2.AuxInt = c
				v1.AddArg2(v2, x)
				v.AddArg2(v0, v1)
				return true
			}
		}
		break
	}
	// match: (Mul64 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
				continue
			}
			v.reset(OpConst64)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (Mul64 (Const64 <t> [c]) (Mul64 (Const64 <t> [d]) x))
	// result: (Mul64 (Const64 <t> [c*d]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpMul64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpMul64)
				v0 := b.NewValue0(v.Pos, OpConst64, t)
				v0.AuxInt = c * d
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpMul64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mul64F (Const64F [c]) (Const64F [d]))
	// result: (Const64F [auxFrom64F(auxTo64F(c) * auxTo64F(d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64F {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64F {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst64F)
			v.AuxInt = auxFrom64F(auxTo64F(c) * auxTo64F(d))
			return true
		}
		break
	}
	// match: (Mul64F x (Const64F [auxFrom64F(1)]))
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpConst64F || v_1.AuxInt != auxFrom64F(1) {
				continue
			}
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Mul64F x (Const64F [auxFrom64F(-1)]))
	// result: (Neg64F x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpConst64F || v_1.AuxInt != auxFrom64F(-1) {
				continue
			}
			v.reset(OpNeg64F)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Mul64F x (Const64F [auxFrom64F(2)]))
	// result: (Add64F x x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpConst64F || v_1.AuxInt != auxFrom64F(2) {
				continue
			}
			v.reset(OpAdd64F)
			v.AddArg2(x, x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpMul8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mul8 (Const8 [c]) (Const8 [d]))
	// result: (Const8 [int64(int8(c*d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst8 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst8)
			v.AuxInt = int64(int8(c * d))
			return true
		}
		break
	}
	// match: (Mul8 (Const8 [1]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 || v_0.AuxInt != 1 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Mul8 (Const8 [-1]) x)
	// result: (Neg8 x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 || v_0.AuxInt != -1 {
				continue
			}
			x := v_1
			v.reset(OpNeg8)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Mul8 <t> n (Const8 [c]))
	// cond: isPowerOfTwo(c)
	// result: (Lsh8x64 <t> n (Const64 <typ.UInt64> [log2(c)]))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpConst8 {
				continue
			}
			c := v_1.AuxInt
			if !(isPowerOfTwo(c)) {
				continue
			}
			v.reset(OpLsh8x64)
			v.Type = t
			v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
			v0.AuxInt = log2(c)
			v.AddArg2(n, v0)
			return true
		}
		break
	}
	// match: (Mul8 <t> n (Const8 [c]))
	// cond: t.IsSigned() && isPowerOfTwo(-c)
	// result: (Neg8 (Lsh8x64 <t> n (Const64 <typ.UInt64> [log2(-c)])))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpConst8 {
				continue
			}
			c := v_1.AuxInt
			if !(t.IsSigned() && isPowerOfTwo(-c)) {
				continue
			}
			v.reset(OpNeg8)
			v0 := b.NewValue0(v.Pos, OpLsh8x64, t)
			v1 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
			v1.AuxInt = log2(-c)
			v0.AddArg2(n, v1)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (Mul8 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
				continue
			}
			v.reset(OpConst8)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (Mul8 (Const8 <t> [c]) (Mul8 (Const8 <t> [d]) x))
	// result: (Mul8 (Const8 <t> [int64(int8(c*d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpMul8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst8 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpMul8)
				v0 := b.NewValue0(v.Pos, OpConst8, t)
				v0.AuxInt = int64(int8(c * d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpNeg16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neg16 (Const16 [c]))
	// result: (Const16 [int64(-int16(c))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(-int16(c))
		return true
	}
	// match: (Neg16 (Sub16 x y))
	// result: (Sub16 y x)
	for {
		if v_0.Op != OpSub16 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpSub16)
		v.AddArg2(y, x)
		return true
	}
	// match: (Neg16 (Neg16 x))
	// result: x
	for {
		if v_0.Op != OpNeg16 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Neg16 <t> (Com16 x))
	// result: (Add16 (Const16 <t> [1]) x)
	for {
		t := v.Type
		if v_0.Op != OpCom16 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = 1
		v.AddArg2(v0, x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeg32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neg32 (Const32 [c]))
	// result: (Const32 [int64(-int32(c))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(-int32(c))
		return true
	}
	// match: (Neg32 (Sub32 x y))
	// result: (Sub32 y x)
	for {
		if v_0.Op != OpSub32 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpSub32)
		v.AddArg2(y, x)
		return true
	}
	// match: (Neg32 (Neg32 x))
	// result: x
	for {
		if v_0.Op != OpNeg32 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Neg32 <t> (Com32 x))
	// result: (Add32 (Const32 <t> [1]) x)
	for {
		t := v.Type
		if v_0.Op != OpCom32 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = 1
		v.AddArg2(v0, x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeg32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg32F (Const32F [c]))
	// cond: auxTo32F(c) != 0
	// result: (Const32F [auxFrom32F(-auxTo32F(c))])
	for {
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		if !(auxTo32F(c) != 0) {
			break
		}
		v.reset(OpConst32F)
		v.AuxInt = auxFrom32F(-auxTo32F(c))
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeg64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neg64 (Const64 [c]))
	// result: (Const64 [-c])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = -c
		return true
	}
	// match: (Neg64 (Sub64 x y))
	// result: (Sub64 y x)
	for {
		if v_0.Op != OpSub64 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpSub64)
		v.AddArg2(y, x)
		return true
	}
	// match: (Neg64 (Neg64 x))
	// result: x
	for {
		if v_0.Op != OpNeg64 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Neg64 <t> (Com64 x))
	// result: (Add64 (Const64 <t> [1]) x)
	for {
		t := v.Type
		if v_0.Op != OpCom64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = 1
		v.AddArg2(v0, x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeg64F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg64F (Const64F [c]))
	// cond: auxTo64F(c) != 0
	// result: (Const64F [auxFrom64F(-auxTo64F(c))])
	for {
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		if !(auxTo64F(c) != 0) {
			break
		}
		v.reset(OpConst64F)
		v.AuxInt = auxFrom64F(-auxTo64F(c))
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeg8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neg8 (Const8 [c]))
	// result: (Const8 [int64( -int8(c))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(-int8(c))
		return true
	}
	// match: (Neg8 (Sub8 x y))
	// result: (Sub8 y x)
	for {
		if v_0.Op != OpSub8 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpSub8)
		v.AddArg2(y, x)
		return true
	}
	// match: (Neg8 (Neg8 x))
	// result: x
	for {
		if v_0.Op != OpNeg8 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Neg8 <t> (Com8 x))
	// result: (Add8 (Const8 <t> [1]) x)
	for {
		t := v.Type
		if v_0.Op != OpCom8 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = 1
		v.AddArg2(v0, x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x x)
	// result: (ConstBool [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}
	// match: (Neq16 (Const16 <t> [c]) (Add16 (Const16 <t> [d]) x))
	// result: (Neq16 (Const16 <t> [int64(int16(c-d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpNeq16)
				v0 := b.NewValue0(v.Pos, OpConst16, t)
				v0.AuxInt = int64(int16(c - d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Neq16 (Const16 [c]) (Const16 [d]))
	// result: (ConstBool [b2i(c != d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst16 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c != d)
			return true
		}
		break
	}
	// match: (Neq16 n (Lsh16x64 (Rsh16x64 (Add16 <t> n (Rsh16Ux64 <t> (Rsh16x64 <t> n (Const64 <typ.UInt64> [15])) (Const64 <typ.UInt64> [kbar]))) (Const64 <typ.UInt64> [k])) (Const64 <typ.UInt64> [k])) )
	// cond: k > 0 && k < 15 && kbar == 16 - k
	// result: (Neq16 (And16 <t> n (Const16 <t> [int64(1<<uint(k)-1)])) (Const16 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpLsh16x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpRsh16x64 {
				continue
			}
			_ = v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAdd16 {
				continue
			}
			t := v_1_0_0.Type
			_ = v_1_0_0.Args[1]
			v_1_0_0_0 := v_1_0_0.Args[0]
			v_1_0_0_1 := v_1_0_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0_0_0, v_1_0_0_1 = _i1+1, v_1_0_0_1, v_1_0_0_0 {
				if n != v_1_0_0_0 || v_1_0_0_1.Op != OpRsh16Ux64 || v_1_0_0_1.Type != t {
					continue
				}
				_ = v_1_0_0_1.Args[1]
				v_1_0_0_1_0 := v_1_0_0_1.Args[0]
				if v_1_0_0_1_0.Op != OpRsh16x64 || v_1_0_0_1_0.Type != t {
					continue
				}
				_ = v_1_0_0_1_0.Args[1]
				if n != v_1_0_0_1_0.Args[0] {
					continue
				}
				v_1_0_0_1_0_1 := v_1_0_0_1_0.Args[1]
				if v_1_0_0_1_0_1.Op != OpConst64 || v_1_0_0_1_0_1.Type != typ.UInt64 || v_1_0_0_1_0_1.AuxInt != 15 {
					continue
				}
				v_1_0_0_1_1 := v_1_0_0_1.Args[1]
				if v_1_0_0_1_1.Op != OpConst64 || v_1_0_0_1_1.Type != typ.UInt64 {
					continue
				}
				kbar := v_1_0_0_1_1.AuxInt
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpConst64 || v_1_0_1.Type != typ.UInt64 {
					continue
				}
				k := v_1_0_1.AuxInt
				v_1_1 := v_1.Args[1]
				if v_1_1.Op != OpConst64 || v_1_1.Type != typ.UInt64 || v_1_1.AuxInt != k || !(k > 0 && k < 15 && kbar == 16-k) {
					continue
				}
				v.reset(OpNeq16)
				v0 := b.NewValue0(v.Pos, OpAnd16, t)
				v1 := b.NewValue0(v.Pos, OpConst16, t)
				v1.AuxInt = int64(1<<uint(k) - 1)
				v0.AddArg2(n, v1)
				v2 := b.NewValue0(v.Pos, OpConst16, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	// match: (Neq16 s:(Sub16 x y) (Const16 [0]))
	// cond: s.Uses == 1
	// result: (Neq16 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s := v_0
			if s.Op != OpSub16 {
				continue
			}
			y := s.Args[1]
			x := s.Args[0]
			if v_1.Op != OpConst16 || v_1.AuxInt != 0 || !(s.Uses == 1) {
				continue
			}
			v.reset(OpNeq16)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (Neq16 (And16 <t> x (Const16 <t> [y])) (Const16 <t> [y]))
	// cond: isPowerOfTwo(y)
	// result: (Eq16 (And16 <t> x (Const16 <t> [y])) (Const16 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd16 {
				continue
			}
			t := v_0.Type
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst16 || v_0_1.Type != t {
					continue
				}
				y := v_0_1.AuxInt
				if v_1.Op != OpConst16 || v_1.Type != t || v_1.AuxInt != y || !(isPowerOfTwo(y)) {
					continue
				}
				v.reset(OpEq16)
				v0 := b.NewValue0(v.Pos, OpAnd16, t)
				v1 := b.NewValue0(v.Pos, OpConst16, t)
				v1.AuxInt = y
				v0.AddArg2(x, v1)
				v2 := b.NewValue0(v.Pos, OpConst16, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq32 x x)
	// result: (ConstBool [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}
	// match: (Neq32 (Const32 <t> [c]) (Add32 (Const32 <t> [d]) x))
	// result: (Neq32 (Const32 <t> [int64(int32(c-d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpNeq32)
				v0 := b.NewValue0(v.Pos, OpConst32, t)
				v0.AuxInt = int64(int32(c - d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Neq32 (Const32 [c]) (Const32 [d]))
	// result: (ConstBool [b2i(c != d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c != d)
			return true
		}
		break
	}
	// match: (Neq32 n (Lsh32x64 (Rsh32x64 (Add32 <t> n (Rsh32Ux64 <t> (Rsh32x64 <t> n (Const64 <typ.UInt64> [31])) (Const64 <typ.UInt64> [kbar]))) (Const64 <typ.UInt64> [k])) (Const64 <typ.UInt64> [k])) )
	// cond: k > 0 && k < 31 && kbar == 32 - k
	// result: (Neq32 (And32 <t> n (Const32 <t> [int64(1<<uint(k)-1)])) (Const32 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpLsh32x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpRsh32x64 {
				continue
			}
			_ = v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAdd32 {
				continue
			}
			t := v_1_0_0.Type
			_ = v_1_0_0.Args[1]
			v_1_0_0_0 := v_1_0_0.Args[0]
			v_1_0_0_1 := v_1_0_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0_0_0, v_1_0_0_1 = _i1+1, v_1_0_0_1, v_1_0_0_0 {
				if n != v_1_0_0_0 || v_1_0_0_1.Op != OpRsh32Ux64 || v_1_0_0_1.Type != t {
					continue
				}
				_ = v_1_0_0_1.Args[1]
				v_1_0_0_1_0 := v_1_0_0_1.Args[0]
				if v_1_0_0_1_0.Op != OpRsh32x64 || v_1_0_0_1_0.Type != t {
					continue
				}
				_ = v_1_0_0_1_0.Args[1]
				if n != v_1_0_0_1_0.Args[0] {
					continue
				}
				v_1_0_0_1_0_1 := v_1_0_0_1_0.Args[1]
				if v_1_0_0_1_0_1.Op != OpConst64 || v_1_0_0_1_0_1.Type != typ.UInt64 || v_1_0_0_1_0_1.AuxInt != 31 {
					continue
				}
				v_1_0_0_1_1 := v_1_0_0_1.Args[1]
				if v_1_0_0_1_1.Op != OpConst64 || v_1_0_0_1_1.Type != typ.UInt64 {
					continue
				}
				kbar := v_1_0_0_1_1.AuxInt
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpConst64 || v_1_0_1.Type != typ.UInt64 {
					continue
				}
				k := v_1_0_1.AuxInt
				v_1_1 := v_1.Args[1]
				if v_1_1.Op != OpConst64 || v_1_1.Type != typ.UInt64 || v_1_1.AuxInt != k || !(k > 0 && k < 31 && kbar == 32-k) {
					continue
				}
				v.reset(OpNeq32)
				v0 := b.NewValue0(v.Pos, OpAnd32, t)
				v1 := b.NewValue0(v.Pos, OpConst32, t)
				v1.AuxInt = int64(1<<uint(k) - 1)
				v0.AddArg2(n, v1)
				v2 := b.NewValue0(v.Pos, OpConst32, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	// match: (Neq32 s:(Sub32 x y) (Const32 [0]))
	// cond: s.Uses == 1
	// result: (Neq32 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s := v_0
			if s.Op != OpSub32 {
				continue
			}
			y := s.Args[1]
			x := s.Args[0]
			if v_1.Op != OpConst32 || v_1.AuxInt != 0 || !(s.Uses == 1) {
				continue
			}
			v.reset(OpNeq32)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (Neq32 (And32 <t> x (Const32 <t> [y])) (Const32 <t> [y]))
	// cond: isPowerOfTwo(y)
	// result: (Eq32 (And32 <t> x (Const32 <t> [y])) (Const32 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd32 {
				continue
			}
			t := v_0.Type
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst32 || v_0_1.Type != t {
					continue
				}
				y := v_0_1.AuxInt
				if v_1.Op != OpConst32 || v_1.Type != t || v_1.AuxInt != y || !(isPowerOfTwo(y)) {
					continue
				}
				v.reset(OpEq32)
				v0 := b.NewValue0(v.Pos, OpAnd32, t)
				v1 := b.NewValue0(v.Pos, OpConst32, t)
				v1.AuxInt = y
				v0.AddArg2(x, v1)
				v2 := b.NewValue0(v.Pos, OpConst32, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpNeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Neq32F (Const32F [c]) (Const32F [d]))
	// result: (ConstBool [b2i(auxTo32F(c) != auxTo32F(d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32F {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32F {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(auxTo32F(c) != auxTo32F(d))
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpNeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq64 x x)
	// result: (ConstBool [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}
	// match: (Neq64 (Const64 <t> [c]) (Add64 (Const64 <t> [d]) x))
	// result: (Neq64 (Const64 <t> [c-d]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpNeq64)
				v0 := b.NewValue0(v.Pos, OpConst64, t)
				v0.AuxInt = c - d
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Neq64 (Const64 [c]) (Const64 [d]))
	// result: (ConstBool [b2i(c != d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c != d)
			return true
		}
		break
	}
	// match: (Neq64 n (Lsh64x64 (Rsh64x64 (Add64 <t> n (Rsh64Ux64 <t> (Rsh64x64 <t> n (Const64 <typ.UInt64> [63])) (Const64 <typ.UInt64> [kbar]))) (Const64 <typ.UInt64> [k])) (Const64 <typ.UInt64> [k])) )
	// cond: k > 0 && k < 63 && kbar == 64 - k
	// result: (Neq64 (And64 <t> n (Const64 <t> [int64(1<<uint(k)-1)])) (Const64 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpLsh64x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpRsh64x64 {
				continue
			}
			_ = v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAdd64 {
				continue
			}
			t := v_1_0_0.Type
			_ = v_1_0_0.Args[1]
			v_1_0_0_0 := v_1_0_0.Args[0]
			v_1_0_0_1 := v_1_0_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0_0_0, v_1_0_0_1 = _i1+1, v_1_0_0_1, v_1_0_0_0 {
				if n != v_1_0_0_0 || v_1_0_0_1.Op != OpRsh64Ux64 || v_1_0_0_1.Type != t {
					continue
				}
				_ = v_1_0_0_1.Args[1]
				v_1_0_0_1_0 := v_1_0_0_1.Args[0]
				if v_1_0_0_1_0.Op != OpRsh64x64 || v_1_0_0_1_0.Type != t {
					continue
				}
				_ = v_1_0_0_1_0.Args[1]
				if n != v_1_0_0_1_0.Args[0] {
					continue
				}
				v_1_0_0_1_0_1 := v_1_0_0_1_0.Args[1]
				if v_1_0_0_1_0_1.Op != OpConst64 || v_1_0_0_1_0_1.Type != typ.UInt64 || v_1_0_0_1_0_1.AuxInt != 63 {
					continue
				}
				v_1_0_0_1_1 := v_1_0_0_1.Args[1]
				if v_1_0_0_1_1.Op != OpConst64 || v_1_0_0_1_1.Type != typ.UInt64 {
					continue
				}
				kbar := v_1_0_0_1_1.AuxInt
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpConst64 || v_1_0_1.Type != typ.UInt64 {
					continue
				}
				k := v_1_0_1.AuxInt
				v_1_1 := v_1.Args[1]
				if v_1_1.Op != OpConst64 || v_1_1.Type != typ.UInt64 || v_1_1.AuxInt != k || !(k > 0 && k < 63 && kbar == 64-k) {
					continue
				}
				v.reset(OpNeq64)
				v0 := b.NewValue0(v.Pos, OpAnd64, t)
				v1 := b.NewValue0(v.Pos, OpConst64, t)
				v1.AuxInt = int64(1<<uint(k) - 1)
				v0.AddArg2(n, v1)
				v2 := b.NewValue0(v.Pos, OpConst64, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	// match: (Neq64 s:(Sub64 x y) (Const64 [0]))
	// cond: s.Uses == 1
	// result: (Neq64 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s := v_0
			if s.Op != OpSub64 {
				continue
			}
			y := s.Args[1]
			x := s.Args[0]
			if v_1.Op != OpConst64 || v_1.AuxInt != 0 || !(s.Uses == 1) {
				continue
			}
			v.reset(OpNeq64)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (Neq64 (And64 <t> x (Const64 <t> [y])) (Const64 <t> [y]))
	// cond: isPowerOfTwo(y)
	// result: (Eq64 (And64 <t> x (Const64 <t> [y])) (Const64 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd64 {
				continue
			}
			t := v_0.Type
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst64 || v_0_1.Type != t {
					continue
				}
				y := v_0_1.AuxInt
				if v_1.Op != OpConst64 || v_1.Type != t || v_1.AuxInt != y || !(isPowerOfTwo(y)) {
					continue
				}
				v.reset(OpEq64)
				v0 := b.NewValue0(v.Pos, OpAnd64, t)
				v1 := b.NewValue0(v.Pos, OpConst64, t)
				v1.AuxInt = y
				v0.AddArg2(x, v1)
				v2 := b.NewValue0(v.Pos, OpConst64, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpNeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Neq64F (Const64F [c]) (Const64F [d]))
	// result: (ConstBool [b2i(auxTo64F(c) != auxTo64F(d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64F {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64F {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(auxTo64F(c) != auxTo64F(d))
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x x)
	// result: (ConstBool [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}
	// match: (Neq8 (Const8 <t> [c]) (Add8 (Const8 <t> [d]) x))
	// result: (Neq8 (Const8 <t> [int64(int8(c-d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpAdd8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst8 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpNeq8)
				v0 := b.NewValue0(v.Pos, OpConst8, t)
				v0.AuxInt = int64(int8(c - d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Neq8 (Const8 [c]) (Const8 [d]))
	// result: (ConstBool [b2i(c != d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst8 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c != d)
			return true
		}
		break
	}
	// match: (Neq8 n (Lsh8x64 (Rsh8x64 (Add8 <t> n (Rsh8Ux64 <t> (Rsh8x64 <t> n (Const64 <typ.UInt64> [ 7])) (Const64 <typ.UInt64> [kbar]))) (Const64 <typ.UInt64> [k])) (Const64 <typ.UInt64> [k])) )
	// cond: k > 0 && k < 7 && kbar == 8 - k
	// result: (Neq8 (And8 <t> n (Const8 <t> [int64(1<<uint(k)-1)])) (Const8 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			n := v_0
			if v_1.Op != OpLsh8x64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpRsh8x64 {
				continue
			}
			_ = v_1_0.Args[1]
			v_1_0_0 := v_1_0.Args[0]
			if v_1_0_0.Op != OpAdd8 {
				continue
			}
			t := v_1_0_0.Type
			_ = v_1_0_0.Args[1]
			v_1_0_0_0 := v_1_0_0.Args[0]
			v_1_0_0_1 := v_1_0_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0_0_0, v_1_0_0_1 = _i1+1, v_1_0_0_1, v_1_0_0_0 {
				if n != v_1_0_0_0 || v_1_0_0_1.Op != OpRsh8Ux64 || v_1_0_0_1.Type != t {
					continue
				}
				_ = v_1_0_0_1.Args[1]
				v_1_0_0_1_0 := v_1_0_0_1.Args[0]
				if v_1_0_0_1_0.Op != OpRsh8x64 || v_1_0_0_1_0.Type != t {
					continue
				}
				_ = v_1_0_0_1_0.Args[1]
				if n != v_1_0_0_1_0.Args[0] {
					continue
				}
				v_1_0_0_1_0_1 := v_1_0_0_1_0.Args[1]
				if v_1_0_0_1_0_1.Op != OpConst64 || v_1_0_0_1_0_1.Type != typ.UInt64 || v_1_0_0_1_0_1.AuxInt != 7 {
					continue
				}
				v_1_0_0_1_1 := v_1_0_0_1.Args[1]
				if v_1_0_0_1_1.Op != OpConst64 || v_1_0_0_1_1.Type != typ.UInt64 {
					continue
				}
				kbar := v_1_0_0_1_1.AuxInt
				v_1_0_1 := v_1_0.Args[1]
				if v_1_0_1.Op != OpConst64 || v_1_0_1.Type != typ.UInt64 {
					continue
				}
				k := v_1_0_1.AuxInt
				v_1_1 := v_1.Args[1]
				if v_1_1.Op != OpConst64 || v_1_1.Type != typ.UInt64 || v_1_1.AuxInt != k || !(k > 0 && k < 7 && kbar == 8-k) {
					continue
				}
				v.reset(OpNeq8)
				v0 := b.NewValue0(v.Pos, OpAnd8, t)
				v1 := b.NewValue0(v.Pos, OpConst8, t)
				v1.AuxInt = int64(1<<uint(k) - 1)
				v0.AddArg2(n, v1)
				v2 := b.NewValue0(v.Pos, OpConst8, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	// match: (Neq8 s:(Sub8 x y) (Const8 [0]))
	// cond: s.Uses == 1
	// result: (Neq8 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			s := v_0
			if s.Op != OpSub8 {
				continue
			}
			y := s.Args[1]
			x := s.Args[0]
			if v_1.Op != OpConst8 || v_1.AuxInt != 0 || !(s.Uses == 1) {
				continue
			}
			v.reset(OpNeq8)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	// match: (Neq8 (And8 <t> x (Const8 <t> [y])) (Const8 <t> [y]))
	// cond: isPowerOfTwo(y)
	// result: (Eq8 (And8 <t> x (Const8 <t> [y])) (Const8 <t> [0]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd8 {
				continue
			}
			t := v_0.Type
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst8 || v_0_1.Type != t {
					continue
				}
				y := v_0_1.AuxInt
				if v_1.Op != OpConst8 || v_1.Type != t || v_1.AuxInt != y || !(isPowerOfTwo(y)) {
					continue
				}
				v.reset(OpEq8)
				v0 := b.NewValue0(v.Pos, OpAnd8, t)
				v1 := b.NewValue0(v.Pos, OpConst8, t)
				v1.AuxInt = y
				v0.AddArg2(x, v1)
				v2 := b.NewValue0(v.Pos, OpConst8, t)
				v2.AuxInt = 0
				v.AddArg2(v0, v2)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpNeqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (NeqB (ConstBool [c]) (ConstBool [d]))
	// result: (ConstBool [b2i(c != d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConstBool {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConstBool {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c != d)
			return true
		}
		break
	}
	// match: (NeqB (ConstBool [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConstBool || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (NeqB (ConstBool [1]) x)
	// result: (Not x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConstBool || v_0.AuxInt != 1 {
				continue
			}
			x := v_1
			v.reset(OpNot)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (NeqB (Not x) (Not y))
	// result: (NeqB x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpNot {
				continue
			}
			x := v_0.Args[0]
			if v_1.Op != OpNot {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpNeqB)
			v.AddArg2(x, y)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpNeqInter(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NeqInter x y)
	// result: (NeqPtr (ITab x) (ITab y))
	for {
		x := v_0
		y := v_1
		v.reset(OpNeqPtr)
		v0 := b.NewValue0(v.Pos, OpITab, typ.Uintptr)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpITab, typ.Uintptr)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValuegeneric_OpNeqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (NeqPtr x x)
	// result: (ConstBool [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConstBool)
		v.AuxInt = 0
		return true
	}
	// match: (NeqPtr (Addr {a} _) (Addr {b} _))
	// result: (ConstBool [b2i(a != b)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAddr {
				continue
			}
			a := v_0.Aux
			if v_1.Op != OpAddr {
				continue
			}
			b := v_1.Aux
			v.reset(OpConstBool)
			v.AuxInt = b2i(a != b)
			return true
		}
		break
	}
	// match: (NeqPtr (Addr {a} _) (OffPtr [o] (Addr {b} _)))
	// result: (ConstBool [b2i(a != b || o != 0)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAddr {
				continue
			}
			a := v_0.Aux
			if v_1.Op != OpOffPtr {
				continue
			}
			o := v_1.AuxInt
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpAddr {
				continue
			}
			b := v_1_0.Aux
			v.reset(OpConstBool)
			v.AuxInt = b2i(a != b || o != 0)
			return true
		}
		break
	}
	// match: (NeqPtr (OffPtr [o1] (Addr {a} _)) (OffPtr [o2] (Addr {b} _)))
	// result: (ConstBool [b2i(a != b || o1 != o2)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			o1 := v_0.AuxInt
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpAddr {
				continue
			}
			a := v_0_0.Aux
			if v_1.Op != OpOffPtr {
				continue
			}
			o2 := v_1.AuxInt
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpAddr {
				continue
			}
			b := v_1_0.Aux
			v.reset(OpConstBool)
			v.AuxInt = b2i(a != b || o1 != o2)
			return true
		}
		break
	}
	// match: (NeqPtr (LocalAddr {a} _ _) (LocalAddr {b} _ _))
	// result: (ConstBool [b2i(a != b)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLocalAddr {
				continue
			}
			a := v_0.Aux
			_ = v_0.Args[1]
			if v_1.Op != OpLocalAddr {
				continue
			}
			b := v_1.Aux
			_ = v_1.Args[1]
			v.reset(OpConstBool)
			v.AuxInt = b2i(a != b)
			return true
		}
		break
	}
	// match: (NeqPtr (LocalAddr {a} _ _) (OffPtr [o] (LocalAddr {b} _ _)))
	// result: (ConstBool [b2i(a != b || o != 0)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLocalAddr {
				continue
			}
			a := v_0.Aux
			_ = v_0.Args[1]
			if v_1.Op != OpOffPtr {
				continue
			}
			o := v_1.AuxInt
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpLocalAddr {
				continue
			}
			b := v_1_0.Aux
			_ = v_1_0.Args[1]
			v.reset(OpConstBool)
			v.AuxInt = b2i(a != b || o != 0)
			return true
		}
		break
	}
	// match: (NeqPtr (OffPtr [o1] (LocalAddr {a} _ _)) (OffPtr [o2] (LocalAddr {b} _ _)))
	// result: (ConstBool [b2i(a != b || o1 != o2)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			o1 := v_0.AuxInt
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpLocalAddr {
				continue
			}
			a := v_0_0.Aux
			_ = v_0_0.Args[1]
			if v_1.Op != OpOffPtr {
				continue
			}
			o2 := v_1.AuxInt
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpLocalAddr {
				continue
			}
			b := v_1_0.Aux
			_ = v_1_0.Args[1]
			v.reset(OpConstBool)
			v.AuxInt = b2i(a != b || o1 != o2)
			return true
		}
		break
	}
	// match: (NeqPtr (OffPtr [o1] p1) p2)
	// cond: isSamePtr(p1, p2)
	// result: (ConstBool [b2i(o1 != 0)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			o1 := v_0.AuxInt
			p1 := v_0.Args[0]
			p2 := v_1
			if !(isSamePtr(p1, p2)) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = b2i(o1 != 0)
			return true
		}
		break
	}
	// match: (NeqPtr (OffPtr [o1] p1) (OffPtr [o2] p2))
	// cond: isSamePtr(p1, p2)
	// result: (ConstBool [b2i(o1 != o2)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			o1 := v_0.AuxInt
			p1 := v_0.Args[0]
			if v_1.Op != OpOffPtr {
				continue
			}
			o2 := v_1.AuxInt
			p2 := v_1.Args[0]
			if !(isSamePtr(p1, p2)) {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = b2i(o1 != o2)
			return true
		}
		break
	}
	// match: (NeqPtr (Const32 [c]) (Const32 [d]))
	// result: (ConstBool [b2i(c != d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c != d)
			return true
		}
		break
	}
	// match: (NeqPtr (Const64 [c]) (Const64 [d]))
	// result: (ConstBool [b2i(c != d)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConstBool)
			v.AuxInt = b2i(c != d)
			return true
		}
		break
	}
	// match: (NeqPtr (LocalAddr _ _) (Addr _))
	// result: (ConstBool [1])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLocalAddr {
				continue
			}
			_ = v_0.Args[1]
			if v_1.Op != OpAddr {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (NeqPtr (OffPtr (LocalAddr _ _)) (Addr _))
	// result: (ConstBool [1])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpLocalAddr {
				continue
			}
			_ = v_0_0.Args[1]
			if v_1.Op != OpAddr {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (NeqPtr (LocalAddr _ _) (OffPtr (Addr _)))
	// result: (ConstBool [1])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpLocalAddr {
				continue
			}
			_ = v_0.Args[1]
			if v_1.Op != OpOffPtr {
				continue
			}
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpAddr {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (NeqPtr (OffPtr (LocalAddr _ _)) (OffPtr (Addr _)))
	// result: (ConstBool [1])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOffPtr {
				continue
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpLocalAddr {
				continue
			}
			_ = v_0_0.Args[1]
			if v_1.Op != OpOffPtr {
				continue
			}
			v_1_0 := v_1.Args[0]
			if v_1_0.Op != OpAddr {
				continue
			}
			v.reset(OpConstBool)
			v.AuxInt = 1
			return true
		}
		break
	}
	// match: (NeqPtr (AddPtr p1 o1) p2)
	// cond: isSamePtr(p1, p2)
	// result: (IsNonNil o1)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAddPtr {
				continue
			}
			o1 := v_0.Args[1]
			p1 := v_0.Args[0]
			p2 := v_1
			if !(isSamePtr(p1, p2)) {
				continue
			}
			v.reset(OpIsNonNil)
			v.AddArg(o1)
			return true
		}
		break
	}
	// match: (NeqPtr (Const32 [0]) p)
	// result: (IsNonNil p)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
				continue
			}
			p := v_1
			v.reset(OpIsNonNil)
			v.AddArg(p)
			return true
		}
		break
	}
	// match: (NeqPtr (Const64 [0]) p)
	// result: (IsNonNil p)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
				continue
			}
			p := v_1
			v.reset(OpIsNonNil)
			v.AddArg(p)
			return true
		}
		break
	}
	// match: (NeqPtr (ConstNil) p)
	// result: (IsNonNil p)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConstNil {
				continue
			}
			p := v_1
			v.reset(OpIsNonNil)
			v.AddArg(p)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpNeqSlice(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NeqSlice x y)
	// result: (NeqPtr (SlicePtr x) (SlicePtr y))
	for {
		x := v_0
		y := v_1
		v.reset(OpNeqPtr)
		v0 := b.NewValue0(v.Pos, OpSlicePtr, typ.BytePtr)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSlicePtr, typ.BytePtr)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValuegeneric_OpNilCheck(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	fe := b.Func.fe
	// match: (NilCheck (GetG mem) mem)
	// result: mem
	for {
		if v_0.Op != OpGetG {
			break
		}
		mem := v_0.Args[0]
		if mem != v_1 {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (NilCheck (Load (OffPtr [c] (SP)) (StaticCall {sym} _)) _)
	// cond: isSameSym(sym, "runtime.newobject") && c == config.ctxt.FixedFrameSize() + config.RegSize && warnRule(fe.Debug_checknil(), v, "removed nil check")
	// result: (Invalid)
	for {
		if v_0.Op != OpLoad {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpOffPtr {
			break
		}
		c := v_0_0.AuxInt
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpSP {
			break
		}
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpStaticCall {
			break
		}
		sym := v_0_1.Aux
		if !(isSameSym(sym, "runtime.newobject") && c == config.ctxt.FixedFrameSize()+config.RegSize && warnRule(fe.Debug_checknil(), v, "removed nil check")) {
			break
		}
		v.reset(OpInvalid)
		return true
	}
	// match: (NilCheck (OffPtr (Load (OffPtr [c] (SP)) (StaticCall {sym} _))) _)
	// cond: isSameSym(sym, "runtime.newobject") && c == config.ctxt.FixedFrameSize() + config.RegSize && warnRule(fe.Debug_checknil(), v, "removed nil check")
	// result: (Invalid)
	for {
		if v_0.Op != OpOffPtr {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLoad {
			break
		}
		_ = v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpOffPtr {
			break
		}
		c := v_0_0_0.AuxInt
		v_0_0_0_0 := v_0_0_0.Args[0]
		if v_0_0_0_0.Op != OpSP {
			break
		}
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpStaticCall {
			break
		}
		sym := v_0_0_1.Aux
		if !(isSameSym(sym, "runtime.newobject") && c == config.ctxt.FixedFrameSize()+config.RegSize && warnRule(fe.Debug_checknil(), v, "removed nil check")) {
			break
		}
		v.reset(OpInvalid)
		return true
	}
	return false
}
func rewriteValuegeneric_OpNot(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Not (ConstBool [c]))
	// result: (ConstBool [1-c])
	for {
		if v_0.Op != OpConstBool {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConstBool)
		v.AuxInt = 1 - c
		return true
	}
	// match: (Not (Eq64 x y))
	// result: (Neq64 x y)
	for {
		if v_0.Op != OpEq64 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpNeq64)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Eq32 x y))
	// result: (Neq32 x y)
	for {
		if v_0.Op != OpEq32 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpNeq32)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Eq16 x y))
	// result: (Neq16 x y)
	for {
		if v_0.Op != OpEq16 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpNeq16)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Eq8 x y))
	// result: (Neq8 x y)
	for {
		if v_0.Op != OpEq8 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpNeq8)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (EqB x y))
	// result: (NeqB x y)
	for {
		if v_0.Op != OpEqB {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpNeqB)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (EqPtr x y))
	// result: (NeqPtr x y)
	for {
		if v_0.Op != OpEqPtr {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpNeqPtr)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Eq64F x y))
	// result: (Neq64F x y)
	for {
		if v_0.Op != OpEq64F {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpNeq64F)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Eq32F x y))
	// result: (Neq32F x y)
	for {
		if v_0.Op != OpEq32F {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpNeq32F)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Neq64 x y))
	// result: (Eq64 x y)
	for {
		if v_0.Op != OpNeq64 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpEq64)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Neq32 x y))
	// result: (Eq32 x y)
	for {
		if v_0.Op != OpNeq32 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpEq32)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Neq16 x y))
	// result: (Eq16 x y)
	for {
		if v_0.Op != OpNeq16 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpEq16)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Neq8 x y))
	// result: (Eq8 x y)
	for {
		if v_0.Op != OpNeq8 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpEq8)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (NeqB x y))
	// result: (EqB x y)
	for {
		if v_0.Op != OpNeqB {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpEqB)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (NeqPtr x y))
	// result: (EqPtr x y)
	for {
		if v_0.Op != OpNeqPtr {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpEqPtr)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Neq64F x y))
	// result: (Eq64F x y)
	for {
		if v_0.Op != OpNeq64F {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpEq64F)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Neq32F x y))
	// result: (Eq32F x y)
	for {
		if v_0.Op != OpNeq32F {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpEq32F)
		v.AddArg2(x, y)
		return true
	}
	// match: (Not (Less64 x y))
	// result: (Leq64 y x)
	for {
		if v_0.Op != OpLess64 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLeq64)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Less32 x y))
	// result: (Leq32 y x)
	for {
		if v_0.Op != OpLess32 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLeq32)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Less16 x y))
	// result: (Leq16 y x)
	for {
		if v_0.Op != OpLess16 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLeq16)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Less8 x y))
	// result: (Leq8 y x)
	for {
		if v_0.Op != OpLess8 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLeq8)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Less64U x y))
	// result: (Leq64U y x)
	for {
		if v_0.Op != OpLess64U {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLeq64U)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Less32U x y))
	// result: (Leq32U y x)
	for {
		if v_0.Op != OpLess32U {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLeq32U)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Less16U x y))
	// result: (Leq16U y x)
	for {
		if v_0.Op != OpLess16U {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLeq16U)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Less8U x y))
	// result: (Leq8U y x)
	for {
		if v_0.Op != OpLess8U {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLeq8U)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Leq64 x y))
	// result: (Less64 y x)
	for {
		if v_0.Op != OpLeq64 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLess64)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Leq32 x y))
	// result: (Less32 y x)
	for {
		if v_0.Op != OpLeq32 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLess32)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Leq16 x y))
	// result: (Less16 y x)
	for {
		if v_0.Op != OpLeq16 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLess16)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Leq8 x y))
	// result: (Less8 y x)
	for {
		if v_0.Op != OpLeq8 {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLess8)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Leq64U x y))
	// result: (Less64U y x)
	for {
		if v_0.Op != OpLeq64U {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLess64U)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Leq32U x y))
	// result: (Less32U y x)
	for {
		if v_0.Op != OpLeq32U {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLess32U)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Leq16U x y))
	// result: (Less16U y x)
	for {
		if v_0.Op != OpLeq16U {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLess16U)
		v.AddArg2(y, x)
		return true
	}
	// match: (Not (Leq8U x y))
	// result: (Less8U y x)
	for {
		if v_0.Op != OpLeq8U {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpLess8U)
		v.AddArg2(y, x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (OffPtr (OffPtr p [b]) [a])
	// result: (OffPtr p [a+b])
	for {
		a := v.AuxInt
		if v_0.Op != OpOffPtr {
			break
		}
		b := v_0.AuxInt
		p := v_0.Args[0]
		v.reset(OpOffPtr)
		v.AuxInt = a + b
		v.AddArg(p)
		return true
	}
	// match: (OffPtr p [0])
	// cond: v.Type.Compare(p.Type) == types.CMPeq
	// result: p
	for {
		if v.AuxInt != 0 {
			break
		}
		p := v_0
		if !(v.Type.Compare(p.Type) == types.CMPeq) {
			break
		}
		v.copyOf(p)
		return true
	}
	return false
}
func rewriteValuegeneric_OpOr16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Or16 (Const16 [c]) (Const16 [d]))
	// result: (Const16 [int64(int16(c|d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst16 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst16)
			v.AuxInt = int64(int16(c | d))
			return true
		}
		break
	}
	// match: (Or16 x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Or16 (Const16 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Or16 (Const16 [-1]) _)
	// result: (Const16 [-1])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 || v_0.AuxInt != -1 {
				continue
			}
			v.reset(OpConst16)
			v.AuxInt = -1
			return true
		}
		break
	}
	// match: (Or16 x (Or16 x y))
	// result: (Or16 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpOr16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.reset(OpOr16)
				v.AddArg2(x, y)
				return true
			}
		}
		break
	}
	// match: (Or16 (And16 x (Const16 [c2])) (Const16 <t> [c1]))
	// cond: ^(c1 | c2) == 0
	// result: (Or16 (Const16 <t> [c1]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd16 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst16 {
					continue
				}
				c2 := v_0_1.AuxInt
				if v_1.Op != OpConst16 {
					continue
				}
				t := v_1.Type
				c1 := v_1.AuxInt
				if !(^(c1 | c2) == 0) {
					continue
				}
				v.reset(OpOr16)
				v0 := b.NewValue0(v.Pos, OpConst16, t)
				v0.AuxInt = c1
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Or16 (Or16 i:(Const16 <t>) z) x)
	// cond: (z.Op != OpConst16 && x.Op != OpConst16)
	// result: (Or16 i (Or16 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOr16 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst16 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst16 && x.Op != OpConst16) {
					continue
				}
				v.reset(OpOr16)
				v0 := b.NewValue0(v.Pos, OpOr16, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Or16 (Const16 <t> [c]) (Or16 (Const16 <t> [d]) x))
	// result: (Or16 (Const16 <t> [int64(int16(c|d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpOr16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpOr16)
				v0 := b.NewValue0(v.Pos, OpConst16, t)
				v0.AuxInt = int64(int16(c | d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpOr32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Or32 (Const32 [c]) (Const32 [d]))
	// result: (Const32 [int64(int32(c|d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst32)
			v.AuxInt = int64(int32(c | d))
			return true
		}
		break
	}
	// match: (Or32 x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Or32 (Const32 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Or32 (Const32 [-1]) _)
	// result: (Const32 [-1])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != -1 {
				continue
			}
			v.reset(OpConst32)
			v.AuxInt = -1
			return true
		}
		break
	}
	// match: (Or32 x (Or32 x y))
	// result: (Or32 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpOr32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.reset(OpOr32)
				v.AddArg2(x, y)
				return true
			}
		}
		break
	}
	// match: (Or32 (And32 x (Const32 [c2])) (Const32 <t> [c1]))
	// cond: ^(c1 | c2) == 0
	// result: (Or32 (Const32 <t> [c1]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd32 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst32 {
					continue
				}
				c2 := v_0_1.AuxInt
				if v_1.Op != OpConst32 {
					continue
				}
				t := v_1.Type
				c1 := v_1.AuxInt
				if !(^(c1 | c2) == 0) {
					continue
				}
				v.reset(OpOr32)
				v0 := b.NewValue0(v.Pos, OpConst32, t)
				v0.AuxInt = c1
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Or32 (Or32 i:(Const32 <t>) z) x)
	// cond: (z.Op != OpConst32 && x.Op != OpConst32)
	// result: (Or32 i (Or32 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOr32 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst32 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst32 && x.Op != OpConst32) {
					continue
				}
				v.reset(OpOr32)
				v0 := b.NewValue0(v.Pos, OpOr32, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Or32 (Const32 <t> [c]) (Or32 (Const32 <t> [d]) x))
	// result: (Or32 (Const32 <t> [int64(int32(c|d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpOr32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpOr32)
				v0 := b.NewValue0(v.Pos, OpConst32, t)
				v0.AuxInt = int64(int32(c | d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpOr64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Or64 (Const64 [c]) (Const64 [d]))
	// result: (Const64 [c|d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst64)
			v.AuxInt = c | d
			return true
		}
		break
	}
	// match: (Or64 x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Or64 (Const64 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Or64 (Const64 [-1]) _)
	// result: (Const64 [-1])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != -1 {
				continue
			}
			v.reset(OpConst64)
			v.AuxInt = -1
			return true
		}
		break
	}
	// match: (Or64 x (Or64 x y))
	// result: (Or64 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpOr64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.reset(OpOr64)
				v.AddArg2(x, y)
				return true
			}
		}
		break
	}
	// match: (Or64 (And64 x (Const64 [c2])) (Const64 <t> [c1]))
	// cond: ^(c1 | c2) == 0
	// result: (Or64 (Const64 <t> [c1]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd64 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst64 {
					continue
				}
				c2 := v_0_1.AuxInt
				if v_1.Op != OpConst64 {
					continue
				}
				t := v_1.Type
				c1 := v_1.AuxInt
				if !(^(c1 | c2) == 0) {
					continue
				}
				v.reset(OpOr64)
				v0 := b.NewValue0(v.Pos, OpConst64, t)
				v0.AuxInt = c1
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Or64 (Or64 i:(Const64 <t>) z) x)
	// cond: (z.Op != OpConst64 && x.Op != OpConst64)
	// result: (Or64 i (Or64 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOr64 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst64 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst64 && x.Op != OpConst64) {
					continue
				}
				v.reset(OpOr64)
				v0 := b.NewValue0(v.Pos, OpOr64, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Or64 (Const64 <t> [c]) (Or64 (Const64 <t> [d]) x))
	// result: (Or64 (Const64 <t> [c|d]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpOr64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpOr64)
				v0 := b.NewValue0(v.Pos, OpConst64, t)
				v0.AuxInt = c | d
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpOr8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Or8 (Const8 [c]) (Const8 [d]))
	// result: (Const8 [int64(int8(c|d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst8 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst8)
			v.AuxInt = int64(int8(c | d))
			return true
		}
		break
	}
	// match: (Or8 x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Or8 (Const8 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Or8 (Const8 [-1]) _)
	// result: (Const8 [-1])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 || v_0.AuxInt != -1 {
				continue
			}
			v.reset(OpConst8)
			v.AuxInt = -1
			return true
		}
		break
	}
	// match: (Or8 x (Or8 x y))
	// result: (Or8 x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpOr8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.reset(OpOr8)
				v.AddArg2(x, y)
				return true
			}
		}
		break
	}
	// match: (Or8 (And8 x (Const8 [c2])) (Const8 <t> [c1]))
	// cond: ^(c1 | c2) == 0
	// result: (Or8 (Const8 <t> [c1]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpAnd8 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				x := v_0_0
				if v_0_1.Op != OpConst8 {
					continue
				}
				c2 := v_0_1.AuxInt
				if v_1.Op != OpConst8 {
					continue
				}
				t := v_1.Type
				c1 := v_1.AuxInt
				if !(^(c1 | c2) == 0) {
					continue
				}
				v.reset(OpOr8)
				v0 := b.NewValue0(v.Pos, OpConst8, t)
				v0.AuxInt = c1
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	// match: (Or8 (Or8 i:(Const8 <t>) z) x)
	// cond: (z.Op != OpConst8 && x.Op != OpConst8)
	// result: (Or8 i (Or8 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpOr8 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst8 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst8 && x.Op != OpConst8) {
					continue
				}
				v.reset(OpOr8)
				v0 := b.NewValue0(v.Pos, OpOr8, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Or8 (Const8 <t> [c]) (Or8 (Const8 <t> [d]) x))
	// result: (Or8 (Const8 <t> [int64(int8(c|d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpOr8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst8 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpOr8)
				v0 := b.NewValue0(v.Pos, OpConst8, t)
				v0.AuxInt = int64(int8(c | d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpPhi(v *Value) bool {
	// match: (Phi (Const8 [c]) (Const8 [c]))
	// result: (Const8 [c])
	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst8 || v_1.AuxInt != c || len(v.Args) != 2 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = c
		return true
	}
	// match: (Phi (Const16 [c]) (Const16 [c]))
	// result: (Const16 [c])
	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst16 || v_1.AuxInt != c || len(v.Args) != 2 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = c
		return true
	}
	// match: (Phi (Const32 [c]) (Const32 [c]))
	// result: (Const32 [c])
	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst32 || v_1.AuxInt != c || len(v.Args) != 2 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = c
		return true
	}
	// match: (Phi (Const64 [c]) (Const64 [c]))
	// result: (Const64 [c])
	for {
		_ = v.Args[1]
		v_0 := v.Args[0]
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v_1 := v.Args[1]
		if v_1.Op != OpConst64 || v_1.AuxInt != c || len(v.Args) != 2 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValuegeneric_OpPtrIndex(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (PtrIndex <t> ptr idx)
	// cond: config.PtrSize == 4
	// result: (AddPtr ptr (Mul32 <typ.Int> idx (Const32 <typ.Int> [t.Elem().Size()])))
	for {
		t := v.Type
		ptr := v_0
		idx := v_1
		if !(config.PtrSize == 4) {
			break
		}
		v.reset(OpAddPtr)
		v0 := b.NewValue0(v.Pos, OpMul32, typ.Int)
		v1 := b.NewValue0(v.Pos, OpConst32, typ.Int)
		v1.AuxInt = t.Elem().Size()
		v0.AddArg2(idx, v1)
		v.AddArg2(ptr, v0)
		return true
	}
	// match: (PtrIndex <t> ptr idx)
	// cond: config.PtrSize == 8
	// result: (AddPtr ptr (Mul64 <typ.Int> idx (Const64 <typ.Int> [t.Elem().Size()])))
	for {
		t := v.Type
		ptr := v_0
		idx := v_1
		if !(config.PtrSize == 8) {
			break
		}
		v.reset(OpAddPtr)
		v0 := b.NewValue0(v.Pos, OpMul64, typ.Int)
		v1 := b.NewValue0(v.Pos, OpConst64, typ.Int)
		v1.AuxInt = t.Elem().Size()
		v0.AddArg2(idx, v1)
		v.AddArg2(ptr, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (RotateLeft16 x (Const16 [c]))
	// cond: c%16 == 0
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		if !(c%16 == 0) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRotateLeft32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (RotateLeft32 x (Const32 [c]))
	// cond: c%32 == 0
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		if !(c%32 == 0) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRotateLeft64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (RotateLeft64 x (Const64 [c]))
	// cond: c%64 == 0
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(c%64 == 0) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (RotateLeft8 x (Const8 [c]))
	// cond: c%8 == 0
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		if !(c%8 == 0) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRound32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Round32F x:(Const32F))
	// result: x
	for {
		x := v_0
		if x.Op != OpConst32F {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRound64F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Round64F x:(Const64F))
	// result: x
	for {
		x := v_0
		if x.Op != OpConst64F {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh16Ux16 <t> x (Const16 [c]))
	// result: (Rsh16Ux64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh16Ux16 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh16Ux32 <t> x (Const32 [c]))
	// result: (Rsh16Ux64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh16Ux32 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 (Const16 [c]) (Const64 [d]))
	// result: (Const16 [int64(int16(uint16(c) >> uint64(d)))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(uint16(c) >> uint64(d)))
		return true
	}
	// match: (Rsh16Ux64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Rsh16Ux64 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh16Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (Const16 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh16Ux64 <t> (Rsh16Ux64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Rsh16Ux64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh16Ux64 (Rsh16x64 x _) (Const64 <t> [15]))
	// result: (Rsh16Ux64 x (Const64 <t> [15]))
	for {
		if v_0.Op != OpRsh16x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		if v_1.Op != OpConst64 {
			break
		}
		t := v_1.Type
		if v_1.AuxInt != 15 {
			break
		}
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = 15
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh16Ux64 (Lsh16x64 (Rsh16Ux64 x (Const64 [c1])) (Const64 [c2])) (Const64 [c3]))
	// cond: uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)
	// result: (Rsh16Ux64 x (Const64 <typ.UInt64> [c1-c2+c3]))
	for {
		if v_0.Op != OpLsh16x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh16Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh16Ux64 (Lsh16x64 x (Const64 [8])) (Const64 [8]))
	// result: (ZeroExt8to16 (Trunc16to8 <typ.UInt8> x))
	for {
		if v_0.Op != OpLsh16x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 8 || v_1.Op != OpConst64 || v_1.AuxInt != 8 {
			break
		}
		v.reset(OpZeroExt8to16)
		v0 := b.NewValue0(v.Pos, OpTrunc16to8, typ.UInt8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh16Ux8 <t> x (Const8 [c]))
	// result: (Rsh16Ux64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh16Ux8 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh16x16 <t> x (Const16 [c]))
	// result: (Rsh16x64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh16x16 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh16x32 <t> x (Const32 [c]))
	// result: (Rsh16x64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh16x32 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 (Const16 [c]) (Const64 [d]))
	// result: (Const16 [int64(int16(c) >> uint64(d))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c) >> uint64(d))
		return true
	}
	// match: (Rsh16x64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Rsh16x64 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh16x64 <t> (Rsh16x64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Rsh16x64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpRsh16x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh16x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh16x64 (Lsh16x64 x (Const64 [8])) (Const64 [8]))
	// result: (SignExt8to16 (Trunc16to8 <typ.Int8> x))
	for {
		if v_0.Op != OpLsh16x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 8 || v_1.Op != OpConst64 || v_1.AuxInt != 8 {
			break
		}
		v.reset(OpSignExt8to16)
		v0 := b.NewValue0(v.Pos, OpTrunc16to8, typ.Int8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh16x8 <t> x (Const8 [c]))
	// result: (Rsh16x64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh16x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh16x8 (Const16 [0]) _)
	// result: (Const16 [0])
	for {
		if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32Ux16 <t> x (Const16 [c]))
	// result: (Rsh32Ux64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32Ux16 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32Ux32 <t> x (Const32 [c]))
	// result: (Rsh32Ux64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32Ux32 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux64 (Const32 [c]) (Const64 [d]))
	// result: (Const32 [int64(int32(uint32(c) >> uint64(d)))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		return true
	}
	// match: (Rsh32Ux64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Rsh32Ux64 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh32Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (Const32 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh32Ux64 <t> (Rsh32Ux64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Rsh32Ux64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpRsh32Ux64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32Ux64 (Rsh32x64 x _) (Const64 <t> [31]))
	// result: (Rsh32Ux64 x (Const64 <t> [31]))
	for {
		if v_0.Op != OpRsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		if v_1.Op != OpConst64 {
			break
		}
		t := v_1.Type
		if v_1.AuxInt != 31 {
			break
		}
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = 31
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32Ux64 (Lsh32x64 (Rsh32Ux64 x (Const64 [c1])) (Const64 [c2])) (Const64 [c3]))
	// cond: uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)
	// result: (Rsh32Ux64 x (Const64 <typ.UInt64> [c1-c2+c3]))
	for {
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh32Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32Ux64 (Lsh32x64 x (Const64 [24])) (Const64 [24]))
	// result: (ZeroExt8to32 (Trunc32to8 <typ.UInt8> x))
	for {
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 24 || v_1.Op != OpConst64 || v_1.AuxInt != 24 {
			break
		}
		v.reset(OpZeroExt8to32)
		v0 := b.NewValue0(v.Pos, OpTrunc32to8, typ.UInt8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32Ux64 (Lsh32x64 x (Const64 [16])) (Const64 [16]))
	// result: (ZeroExt16to32 (Trunc32to16 <typ.UInt16> x))
	for {
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 16 || v_1.Op != OpConst64 || v_1.AuxInt != 16 {
			break
		}
		v.reset(OpZeroExt16to32)
		v0 := b.NewValue0(v.Pos, OpTrunc32to16, typ.UInt16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32Ux8 <t> x (Const8 [c]))
	// result: (Rsh32Ux64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32Ux8 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32x16 <t> x (Const16 [c]))
	// result: (Rsh32x64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32x16 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32x32 <t> x (Const32 [c]))
	// result: (Rsh32x64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32x32 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x64 (Const32 [c]) (Const64 [d]))
	// result: (Const32 [int64(int32(c) >> uint64(d))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c) >> uint64(d))
		return true
	}
	// match: (Rsh32x64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Rsh32x64 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh32x64 <t> (Rsh32x64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Rsh32x64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpRsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh32x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32x64 (Lsh32x64 x (Const64 [24])) (Const64 [24]))
	// result: (SignExt8to32 (Trunc32to8 <typ.Int8> x))
	for {
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 24 || v_1.Op != OpConst64 || v_1.AuxInt != 24 {
			break
		}
		v.reset(OpSignExt8to32)
		v0 := b.NewValue0(v.Pos, OpTrunc32to8, typ.Int8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x64 (Lsh32x64 x (Const64 [16])) (Const64 [16]))
	// result: (SignExt16to32 (Trunc32to16 <typ.Int16> x))
	for {
		if v_0.Op != OpLsh32x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 16 || v_1.Op != OpConst64 || v_1.AuxInt != 16 {
			break
		}
		v.reset(OpSignExt16to32)
		v0 := b.NewValue0(v.Pos, OpTrunc32to16, typ.Int16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32x8 <t> x (Const8 [c]))
	// result: (Rsh32x64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh32x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh32x8 (Const32 [0]) _)
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh64Ux16 <t> x (Const16 [c]))
	// result: (Rsh64Ux64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64Ux16 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh64Ux32 <t> x (Const32 [c]))
	// result: (Rsh64Ux64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64Ux32 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux64 (Const64 [c]) (Const64 [d]))
	// result: (Const64 [int64(uint64(c) >> uint64(d))])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(uint64(c) >> uint64(d))
		return true
	}
	// match: (Rsh64Ux64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Rsh64Ux64 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh64Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (Const64 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh64Ux64 <t> (Rsh64Ux64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Rsh64Ux64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpRsh64Ux64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64Ux64 (Rsh64x64 x _) (Const64 <t> [63]))
	// result: (Rsh64Ux64 x (Const64 <t> [63]))
	for {
		if v_0.Op != OpRsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		if v_1.Op != OpConst64 {
			break
		}
		t := v_1.Type
		if v_1.AuxInt != 63 {
			break
		}
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = 63
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64Ux64 (Lsh64x64 (Rsh64Ux64 x (Const64 [c1])) (Const64 [c2])) (Const64 [c3]))
	// cond: uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)
	// result: (Rsh64Ux64 x (Const64 <typ.UInt64> [c1-c2+c3]))
	for {
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh64Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64Ux64 (Lsh64x64 x (Const64 [56])) (Const64 [56]))
	// result: (ZeroExt8to64 (Trunc64to8 <typ.UInt8> x))
	for {
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 56 || v_1.Op != OpConst64 || v_1.AuxInt != 56 {
			break
		}
		v.reset(OpZeroExt8to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to8, typ.UInt8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64Ux64 (Lsh64x64 x (Const64 [48])) (Const64 [48]))
	// result: (ZeroExt16to64 (Trunc64to16 <typ.UInt16> x))
	for {
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 48 || v_1.Op != OpConst64 || v_1.AuxInt != 48 {
			break
		}
		v.reset(OpZeroExt16to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to16, typ.UInt16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64Ux64 (Lsh64x64 x (Const64 [32])) (Const64 [32]))
	// result: (ZeroExt32to64 (Trunc64to32 <typ.UInt32> x))
	for {
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 32 || v_1.Op != OpConst64 || v_1.AuxInt != 32 {
			break
		}
		v.reset(OpZeroExt32to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh64Ux8 <t> x (Const8 [c]))
	// result: (Rsh64Ux64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64Ux8 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh64x16 <t> x (Const16 [c]))
	// result: (Rsh64x64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64x16 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh64x32 <t> x (Const32 [c]))
	// result: (Rsh64x64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64x32 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x64 (Const64 [c]) (Const64 [d]))
	// result: (Const64 [c >> uint64(d)])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c >> uint64(d)
		return true
	}
	// match: (Rsh64x64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Rsh64x64 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh64x64 <t> (Rsh64x64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Rsh64x64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpRsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64x64 (Lsh64x64 x (Const64 [56])) (Const64 [56]))
	// result: (SignExt8to64 (Trunc64to8 <typ.Int8> x))
	for {
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 56 || v_1.Op != OpConst64 || v_1.AuxInt != 56 {
			break
		}
		v.reset(OpSignExt8to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to8, typ.Int8)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64x64 (Lsh64x64 x (Const64 [48])) (Const64 [48]))
	// result: (SignExt16to64 (Trunc64to16 <typ.Int16> x))
	for {
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 48 || v_1.Op != OpConst64 || v_1.AuxInt != 48 {
			break
		}
		v.reset(OpSignExt16to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to16, typ.Int16)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh64x64 (Lsh64x64 x (Const64 [32])) (Const64 [32]))
	// result: (SignExt32to64 (Trunc64to32 <typ.Int32> x))
	for {
		if v_0.Op != OpLsh64x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 || v_0_1.AuxInt != 32 || v_1.Op != OpConst64 || v_1.AuxInt != 32 {
			break
		}
		v.reset(OpSignExt32to64)
		v0 := b.NewValue0(v.Pos, OpTrunc64to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh64x8 <t> x (Const8 [c]))
	// result: (Rsh64x64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64x8 (Const64 [0]) _)
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh8Ux16 <t> x (Const16 [c]))
	// result: (Rsh8Ux64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh8Ux16 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh8Ux32 <t> x (Const32 [c]))
	// result: (Rsh8Ux64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh8Ux32 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 (Const8 [c]) (Const64 [d]))
	// result: (Const8 [int64(int8(uint8(c) >> uint64(d)))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(uint8(c) >> uint64(d)))
		return true
	}
	// match: (Rsh8Ux64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Rsh8Ux64 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh8Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (Const8 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh8Ux64 <t> (Rsh8Ux64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Rsh8Ux64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh8Ux64 (Rsh8x64 x _) (Const64 <t> [7] ))
	// result: (Rsh8Ux64 x (Const64 <t> [7] ))
	for {
		if v_0.Op != OpRsh8x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		if v_1.Op != OpConst64 {
			break
		}
		t := v_1.Type
		if v_1.AuxInt != 7 {
			break
		}
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = 7
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh8Ux64 (Lsh8x64 (Rsh8Ux64 x (Const64 [c1])) (Const64 [c2])) (Const64 [c3]))
	// cond: uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)
	// result: (Rsh8Ux64 x (Const64 <typ.UInt64> [c1-c2+c3]))
	for {
		if v_0.Op != OpLsh8x64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpRsh8Ux64 {
			break
		}
		_ = v_0_0.Args[1]
		x := v_0_0.Args[0]
		v_0_0_1 := v_0_0.Args[1]
		if v_0_0_1.Op != OpConst64 {
			break
		}
		c1 := v_0_0_1.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c2 := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		c3 := v_1.AuxInt
		if !(uint64(c1) >= uint64(c2) && uint64(c3) >= uint64(c2) && !uaddOvf(c1-c2, c3)) {
			break
		}
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, typ.UInt64)
		v0.AuxInt = c1 - c2 + c3
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh8Ux8 <t> x (Const8 [c]))
	// result: (Rsh8Ux64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh8Ux8 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh8x16 <t> x (Const16 [c]))
	// result: (Rsh8x64 x (Const64 <t> [int64(uint16(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint16(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh8x16 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh8x32 <t> x (Const32 [c]))
	// result: (Rsh8x64 x (Const64 <t> [int64(uint32(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint32(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh8x32 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh8x64 (Const8 [c]) (Const64 [d]))
	// result: (Const8 [int64(int8(c) >> uint64(d))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c) >> uint64(d))
		return true
	}
	// match: (Rsh8x64 x (Const64 [0]))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpConst64 || v_1.AuxInt != 0 {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (Rsh8x64 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	// match: (Rsh8x64 <t> (Rsh8x64 x (Const64 [c])) (Const64 [d]))
	// cond: !uaddOvf(c,d)
	// result: (Rsh8x64 x (Const64 <t> [c+d]))
	for {
		t := v.Type
		if v_0.Op != OpRsh8x64 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		c := v_0_1.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		if !(!uaddOvf(c, d)) {
			break
		}
		v.reset(OpRsh8x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(x, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh8x8 <t> x (Const8 [c]))
	// result: (Rsh8x64 x (Const64 <t> [int64(uint8(c))]))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		c := v_1.AuxInt
		v.reset(OpRsh8x64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = int64(uint8(c))
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh8x8 (Const8 [0]) _)
	// result: (Const8 [0])
	for {
		if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpSelect0(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Select0 (Div128u (Const64 [0]) lo y))
	// result: (Div64u lo y)
	for {
		if v_0.Op != OpDiv128u {
			break
		}
		y := v_0.Args[2]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 || v_0_0.AuxInt != 0 {
			break
		}
		lo := v_0.Args[1]
		v.reset(OpDiv64u)
		v.AddArg2(lo, y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSelect1(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Select1 (Div128u (Const64 [0]) lo y))
	// result: (Mod64u lo y)
	for {
		if v_0.Op != OpDiv128u {
			break
		}
		y := v_0.Args[2]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpConst64 || v_0_0.AuxInt != 0 {
			break
		}
		lo := v_0.Args[1]
		v.reset(OpMod64u)
		v.AddArg2(lo, y)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt16to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SignExt16to32 (Const16 [c]))
	// result: (Const32 [int64( int16(c))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int16(c))
		return true
	}
	// match: (SignExt16to32 (Trunc32to16 x:(Rsh32x64 _ (Const64 [s]))))
	// cond: s >= 16
	// result: x
	for {
		if v_0.Op != OpTrunc32to16 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh32x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 16) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt16to64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SignExt16to64 (Const16 [c]))
	// result: (Const64 [int64( int16(c))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(int16(c))
		return true
	}
	// match: (SignExt16to64 (Trunc64to16 x:(Rsh64x64 _ (Const64 [s]))))
	// cond: s >= 48
	// result: x
	for {
		if v_0.Op != OpTrunc64to16 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 48) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt32to64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SignExt32to64 (Const32 [c]))
	// result: (Const64 [int64( int32(c))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(int32(c))
		return true
	}
	// match: (SignExt32to64 (Trunc64to32 x:(Rsh64x64 _ (Const64 [s]))))
	// cond: s >= 32
	// result: x
	for {
		if v_0.Op != OpTrunc64to32 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 32) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt8to16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SignExt8to16 (Const8 [c]))
	// result: (Const16 [int64( int8(c))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int8(c))
		return true
	}
	// match: (SignExt8to16 (Trunc16to8 x:(Rsh16x64 _ (Const64 [s]))))
	// cond: s >= 8
	// result: x
	for {
		if v_0.Op != OpTrunc16to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh16x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 8) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt8to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SignExt8to32 (Const8 [c]))
	// result: (Const32 [int64( int8(c))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int8(c))
		return true
	}
	// match: (SignExt8to32 (Trunc32to8 x:(Rsh32x64 _ (Const64 [s]))))
	// cond: s >= 24
	// result: x
	for {
		if v_0.Op != OpTrunc32to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh32x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 24) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSignExt8to64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SignExt8to64 (Const8 [c]))
	// result: (Const64 [int64( int8(c))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(int8(c))
		return true
	}
	// match: (SignExt8to64 (Trunc64to8 x:(Rsh64x64 _ (Const64 [s]))))
	// cond: s >= 56
	// result: x
	for {
		if v_0.Op != OpTrunc64to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64x64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 56) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSliceCap(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SliceCap (SliceMake _ _ (Const64 <t> [c])))
	// result: (Const64 <t> [c])
	for {
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_2 := v_0.Args[2]
		if v_0_2.Op != OpConst64 {
			break
		}
		t := v_0_2.Type
		c := v_0_2.AuxInt
		v.reset(OpConst64)
		v.Type = t
		v.AuxInt = c
		return true
	}
	// match: (SliceCap (SliceMake _ _ (Const32 <t> [c])))
	// result: (Const32 <t> [c])
	for {
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_2 := v_0.Args[2]
		if v_0_2.Op != OpConst32 {
			break
		}
		t := v_0_2.Type
		c := v_0_2.AuxInt
		v.reset(OpConst32)
		v.Type = t
		v.AuxInt = c
		return true
	}
	// match: (SliceCap (SliceMake _ _ (SliceCap x)))
	// result: (SliceCap x)
	for {
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_2 := v_0.Args[2]
		if v_0_2.Op != OpSliceCap {
			break
		}
		x := v_0_2.Args[0]
		v.reset(OpSliceCap)
		v.AddArg(x)
		return true
	}
	// match: (SliceCap (SliceMake _ _ (SliceLen x)))
	// result: (SliceLen x)
	for {
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_2 := v_0.Args[2]
		if v_0_2.Op != OpSliceLen {
			break
		}
		x := v_0_2.Args[0]
		v.reset(OpSliceLen)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSliceLen(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SliceLen (SliceMake _ (Const64 <t> [c]) _))
	// result: (Const64 <t> [c])
	for {
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		c := v_0_1.AuxInt
		v.reset(OpConst64)
		v.Type = t
		v.AuxInt = c
		return true
	}
	// match: (SliceLen (SliceMake _ (Const32 <t> [c]) _))
	// result: (Const32 <t> [c])
	for {
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst32 {
			break
		}
		t := v_0_1.Type
		c := v_0_1.AuxInt
		v.reset(OpConst32)
		v.Type = t
		v.AuxInt = c
		return true
	}
	// match: (SliceLen (SliceMake _ (SliceLen x) _))
	// result: (SliceLen x)
	for {
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpSliceLen {
			break
		}
		x := v_0_1.Args[0]
		v.reset(OpSliceLen)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSlicePtr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SlicePtr (SliceMake (SlicePtr x) _ _))
	// result: (SlicePtr x)
	for {
		if v_0.Op != OpSliceMake {
			break
		}
		_ = v_0.Args[2]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpSlicePtr {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpSlicePtr)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Slicemask (Const32 [x]))
	// cond: x > 0
	// result: (Const32 [-1])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		x := v_0.AuxInt
		if !(x > 0) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = -1
		return true
	}
	// match: (Slicemask (Const32 [0]))
	// result: (Const32 [0])
	for {
		if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	// match: (Slicemask (Const64 [x]))
	// cond: x > 0
	// result: (Const64 [-1])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		x := v_0.AuxInt
		if !(x > 0) {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = -1
		return true
	}
	// match: (Slicemask (Const64 [0]))
	// result: (Const64 [0])
	for {
		if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValuegeneric_OpSqrt(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Sqrt (Const64F [c]))
	// result: (Const64F [auxFrom64F(math.Sqrt(auxTo64F(c)))])
	for {
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = auxFrom64F(math.Sqrt(auxTo64F(c)))
		return true
	}
	return false
}
func rewriteValuegeneric_OpStaticCall(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (StaticCall {sym} s1:(Store _ (Const64 [sz]) s2:(Store _ src s3:(Store {t} _ dst mem))))
	// cond: sz >= 0 && isSameSym(sym,"runtime.memmove") && t.(*types.Type).IsPtr() && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && isInlinableMemmove(dst,src,sz,config) && clobber(s1, s2, s3)
	// result: (Move {t.(*types.Type).Elem()} [sz] dst src mem)
	for {
		sym := v.Aux
		s1 := v_0
		if s1.Op != OpStore {
			break
		}
		_ = s1.Args[2]
		s1_1 := s1.Args[1]
		if s1_1.Op != OpConst64 {
			break
		}
		sz := s1_1.AuxInt
		s2 := s1.Args[2]
		if s2.Op != OpStore {
			break
		}
		_ = s2.Args[2]
		src := s2.Args[1]
		s3 := s2.Args[2]
		if s3.Op != OpStore {
			break
		}
		t := s3.Aux
		mem := s3.Args[2]
		dst := s3.Args[1]
		if !(sz >= 0 && isSameSym(sym, "runtime.memmove") && t.(*types.Type).IsPtr() && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && isInlinableMemmove(dst, src, sz, config) && clobber(s1, s2, s3)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = sz
		v.Aux = t.(*types.Type).Elem()
		v.AddArg3(dst, src, mem)
		return true
	}
	// match: (StaticCall {sym} s1:(Store _ (Const32 [sz]) s2:(Store _ src s3:(Store {t} _ dst mem))))
	// cond: sz >= 0 && isSameSym(sym,"runtime.memmove") && t.(*types.Type).IsPtr() && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && isInlinableMemmove(dst,src,sz,config) && clobber(s1, s2, s3)
	// result: (Move {t.(*types.Type).Elem()} [sz] dst src mem)
	for {
		sym := v.Aux
		s1 := v_0
		if s1.Op != OpStore {
			break
		}
		_ = s1.Args[2]
		s1_1 := s1.Args[1]
		if s1_1.Op != OpConst32 {
			break
		}
		sz := s1_1.AuxInt
		s2 := s1.Args[2]
		if s2.Op != OpStore {
			break
		}
		_ = s2.Args[2]
		src := s2.Args[1]
		s3 := s2.Args[2]
		if s3.Op != OpStore {
			break
		}
		t := s3.Aux
		mem := s3.Args[2]
		dst := s3.Args[1]
		if !(sz >= 0 && isSameSym(sym, "runtime.memmove") && t.(*types.Type).IsPtr() && s1.Uses == 1 && s2.Uses == 1 && s3.Uses == 1 && isInlinableMemmove(dst, src, sz, config) && clobber(s1, s2, s3)) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = sz
		v.Aux = t.(*types.Type).Elem()
		v.AddArg3(dst, src, mem)
		return true
	}
	// match: (StaticCall {sym} x)
	// cond: needRaceCleanup(sym,v)
	// result: x
	for {
		sym := v.Aux
		x := v_0
		if !(needRaceCleanup(sym, v)) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	fe := b.Func.fe
	// match: (Store {t1} p1 (Load <t2> p2 mem) mem)
	// cond: isSamePtr(p1, p2) && t2.Size() == sizeof(t1)
	// result: mem
	for {
		t1 := v.Aux
		p1 := v_0
		if v_1.Op != OpLoad {
			break
		}
		t2 := v_1.Type
		mem := v_1.Args[1]
		p2 := v_1.Args[0]
		if mem != v_2 || !(isSamePtr(p1, p2) && t2.Size() == sizeof(t1)) {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (Store {t1} p1 (Load <t2> p2 oldmem) mem:(Store {t3} p3 _ oldmem))
	// cond: isSamePtr(p1, p2) && t2.Size() == sizeof(t1) && disjoint(p1, sizeof(t1), p3, sizeof(t3))
	// result: mem
	for {
		t1 := v.Aux
		p1 := v_0
		if v_1.Op != OpLoad {
			break
		}
		t2 := v_1.Type
		oldmem := v_1.Args[1]
		p2 := v_1.Args[0]
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t3 := mem.Aux
		_ = mem.Args[2]
		p3 := mem.Args[0]
		if oldmem != mem.Args[2] || !(isSamePtr(p1, p2) && t2.Size() == sizeof(t1) && disjoint(p1, sizeof(t1), p3, sizeof(t3))) {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (Store {t1} p1 (Load <t2> p2 oldmem) mem:(Store {t3} p3 _ (Store {t4} p4 _ oldmem)))
	// cond: isSamePtr(p1, p2) && t2.Size() == sizeof(t1) && disjoint(p1, sizeof(t1), p3, sizeof(t3)) && disjoint(p1, sizeof(t1), p4, sizeof(t4))
	// result: mem
	for {
		t1 := v.Aux
		p1 := v_0
		if v_1.Op != OpLoad {
			break
		}
		t2 := v_1.Type
		oldmem := v_1.Args[1]
		p2 := v_1.Args[0]
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t3 := mem.Aux
		_ = mem.Args[2]
		p3 := mem.Args[0]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t4 := mem_2.Aux
		_ = mem_2.Args[2]
		p4 := mem_2.Args[0]
		if oldmem != mem_2.Args[2] || !(isSamePtr(p1, p2) && t2.Size() == sizeof(t1) && disjoint(p1, sizeof(t1), p3, sizeof(t3)) && disjoint(p1, sizeof(t1), p4, sizeof(t4))) {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (Store {t1} p1 (Load <t2> p2 oldmem) mem:(Store {t3} p3 _ (Store {t4} p4 _ (Store {t5} p5 _ oldmem))))
	// cond: isSamePtr(p1, p2) && t2.Size() == sizeof(t1) && disjoint(p1, sizeof(t1), p3, sizeof(t3)) && disjoint(p1, sizeof(t1), p4, sizeof(t4)) && disjoint(p1, sizeof(t1), p5, sizeof(t5))
	// result: mem
	for {
		t1 := v.Aux
		p1 := v_0
		if v_1.Op != OpLoad {
			break
		}
		t2 := v_1.Type
		oldmem := v_1.Args[1]
		p2 := v_1.Args[0]
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t3 := mem.Aux
		_ = mem.Args[2]
		p3 := mem.Args[0]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t4 := mem_2.Aux
		_ = mem_2.Args[2]
		p4 := mem_2.Args[0]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t5 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		p5 := mem_2_2.Args[0]
		if oldmem != mem_2_2.Args[2] || !(isSamePtr(p1, p2) && t2.Size() == sizeof(t1) && disjoint(p1, sizeof(t1), p3, sizeof(t3)) && disjoint(p1, sizeof(t1), p4, sizeof(t4)) && disjoint(p1, sizeof(t1), p5, sizeof(t5))) {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (Store {t} (OffPtr [o] p1) x mem:(Zero [n] p2 _))
	// cond: isConstZero(x) && o >= 0 && sizeof(t) + o <= n && isSamePtr(p1, p2)
	// result: mem
	for {
		t := v.Aux
		if v_0.Op != OpOffPtr {
			break
		}
		o := v_0.AuxInt
		p1 := v_0.Args[0]
		x := v_1
		mem := v_2
		if mem.Op != OpZero {
			break
		}
		n := mem.AuxInt
		_ = mem.Args[1]
		p2 := mem.Args[0]
		if !(isConstZero(x) && o >= 0 && sizeof(t)+o <= n && isSamePtr(p1, p2)) {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (Store {t1} op:(OffPtr [o1] p1) x mem:(Store {t2} p2 _ (Zero [n] p3 _)))
	// cond: isConstZero(x) && o1 >= 0 && sizeof(t1) + o1 <= n && isSamePtr(p1, p3) && disjoint(op, sizeof(t1), p2, sizeof(t2))
	// result: mem
	for {
		t1 := v.Aux
		op := v_0
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		x := v_1
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		p2 := mem.Args[0]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpZero {
			break
		}
		n := mem_2.AuxInt
		_ = mem_2.Args[1]
		p3 := mem_2.Args[0]
		if !(isConstZero(x) && o1 >= 0 && sizeof(t1)+o1 <= n && isSamePtr(p1, p3) && disjoint(op, sizeof(t1), p2, sizeof(t2))) {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (Store {t1} op:(OffPtr [o1] p1) x mem:(Store {t2} p2 _ (Store {t3} p3 _ (Zero [n] p4 _))))
	// cond: isConstZero(x) && o1 >= 0 && sizeof(t1) + o1 <= n && isSamePtr(p1, p4) && disjoint(op, sizeof(t1), p2, sizeof(t2)) && disjoint(op, sizeof(t1), p3, sizeof(t3))
	// result: mem
	for {
		t1 := v.Aux
		op := v_0
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		x := v_1
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		p2 := mem.Args[0]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		p3 := mem_2.Args[0]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpZero {
			break
		}
		n := mem_2_2.AuxInt
		_ = mem_2_2.Args[1]
		p4 := mem_2_2.Args[0]
		if !(isConstZero(x) && o1 >= 0 && sizeof(t1)+o1 <= n && isSamePtr(p1, p4) && disjoint(op, sizeof(t1), p2, sizeof(t2)) && disjoint(op, sizeof(t1), p3, sizeof(t3))) {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (Store {t1} op:(OffPtr [o1] p1) x mem:(Store {t2} p2 _ (Store {t3} p3 _ (Store {t4} p4 _ (Zero [n] p5 _)))))
	// cond: isConstZero(x) && o1 >= 0 && sizeof(t1) + o1 <= n && isSamePtr(p1, p5) && disjoint(op, sizeof(t1), p2, sizeof(t2)) && disjoint(op, sizeof(t1), p3, sizeof(t3)) && disjoint(op, sizeof(t1), p4, sizeof(t4))
	// result: mem
	for {
		t1 := v.Aux
		op := v_0
		if op.Op != OpOffPtr {
			break
		}
		o1 := op.AuxInt
		p1 := op.Args[0]
		x := v_1
		mem := v_2
		if mem.Op != OpStore {
			break
		}
		t2 := mem.Aux
		_ = mem.Args[2]
		p2 := mem.Args[0]
		mem_2 := mem.Args[2]
		if mem_2.Op != OpStore {
			break
		}
		t3 := mem_2.Aux
		_ = mem_2.Args[2]
		p3 := mem_2.Args[0]
		mem_2_2 := mem_2.Args[2]
		if mem_2_2.Op != OpStore {
			break
		}
		t4 := mem_2_2.Aux
		_ = mem_2_2.Args[2]
		p4 := mem_2_2.Args[0]
		mem_2_2_2 := mem_2_2.Args[2]
		if mem_2_2_2.Op != OpZero {
			break
		}
		n := mem_2_2_2.AuxInt
		_ = mem_2_2_2.Args[1]
		p5 := mem_2_2_2.Args[0]
		if !(isConstZero(x) && o1 >= 0 && sizeof(t1)+o1 <= n && isSamePtr(p1, p5) && disjoint(op, sizeof(t1), p2, sizeof(t2)) && disjoint(op, sizeof(t1), p3, sizeof(t3)) && disjoint(op, sizeof(t1), p4, sizeof(t4))) {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (Store _ (StructMake0) mem)
	// result: mem
	for {
		if v_1.Op != OpStructMake0 {
			break
		}
		mem := v_2
		v.copyOf(mem)
		return true
	}
	// match: (Store dst (StructMake1 <t> f0) mem)
	// result: (Store {t.FieldType(0)} (OffPtr <t.FieldType(0).PtrTo()> [0] dst) f0 mem)
	for {
		dst := v_0
		if v_1.Op != OpStructMake1 {
			break
		}
		t := v_1.Type
		f0 := v_1.Args[0]
		mem := v_2
		v.reset(OpStore)
		v.Aux = t.FieldType(0)
		v0 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(0).PtrTo())
		v0.AuxInt = 0
		v0.AddArg(dst)
		v.AddArg3(v0, f0, mem)
		return true
	}
	// match: (Store dst (StructMake2 <t> f0 f1) mem)
	// result: (Store {t.FieldType(1)} (OffPtr <t.FieldType(1).PtrTo()> [t.FieldOff(1)] dst) f1 (Store {t.FieldType(0)} (OffPtr <t.FieldType(0).PtrTo()> [0] dst) f0 mem))
	for {
		dst := v_0
		if v_1.Op != OpStructMake2 {
			break
		}
		t := v_1.Type
		f1 := v_1.Args[1]
		f0 := v_1.Args[0]
		mem := v_2
		v.reset(OpStore)
		v.Aux = t.FieldType(1)
		v0 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(1).PtrTo())
		v0.AuxInt = t.FieldOff(1)
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t.FieldType(0)
		v2 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(0).PtrTo())
		v2.AuxInt = 0
		v2.AddArg(dst)
		v1.AddArg3(v2, f0, mem)
		v.AddArg3(v0, f1, v1)
		return true
	}
	// match: (Store dst (StructMake3 <t> f0 f1 f2) mem)
	// result: (Store {t.FieldType(2)} (OffPtr <t.FieldType(2).PtrTo()> [t.FieldOff(2)] dst) f2 (Store {t.FieldType(1)} (OffPtr <t.FieldType(1).PtrTo()> [t.FieldOff(1)] dst) f1 (Store {t.FieldType(0)} (OffPtr <t.FieldType(0).PtrTo()> [0] dst) f0 mem)))
	for {
		dst := v_0
		if v_1.Op != OpStructMake3 {
			break
		}
		t := v_1.Type
		f2 := v_1.Args[2]
		f0 := v_1.Args[0]
		f1 := v_1.Args[1]
		mem := v_2
		v.reset(OpStore)
		v.Aux = t.FieldType(2)
		v0 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(2).PtrTo())
		v0.AuxInt = t.FieldOff(2)
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t.FieldType(1)
		v2 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(1).PtrTo())
		v2.AuxInt = t.FieldOff(1)
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v3.Aux = t.FieldType(0)
		v4 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(0).PtrTo())
		v4.AuxInt = 0
		v4.AddArg(dst)
		v3.AddArg3(v4, f0, mem)
		v1.AddArg3(v2, f1, v3)
		v.AddArg3(v0, f2, v1)
		return true
	}
	// match: (Store dst (StructMake4 <t> f0 f1 f2 f3) mem)
	// result: (Store {t.FieldType(3)} (OffPtr <t.FieldType(3).PtrTo()> [t.FieldOff(3)] dst) f3 (Store {t.FieldType(2)} (OffPtr <t.FieldType(2).PtrTo()> [t.FieldOff(2)] dst) f2 (Store {t.FieldType(1)} (OffPtr <t.FieldType(1).PtrTo()> [t.FieldOff(1)] dst) f1 (Store {t.FieldType(0)} (OffPtr <t.FieldType(0).PtrTo()> [0] dst) f0 mem))))
	for {
		dst := v_0
		if v_1.Op != OpStructMake4 {
			break
		}
		t := v_1.Type
		f3 := v_1.Args[3]
		f0 := v_1.Args[0]
		f1 := v_1.Args[1]
		f2 := v_1.Args[2]
		mem := v_2
		v.reset(OpStore)
		v.Aux = t.FieldType(3)
		v0 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(3).PtrTo())
		v0.AuxInt = t.FieldOff(3)
		v0.AddArg(dst)
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t.FieldType(2)
		v2 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(2).PtrTo())
		v2.AuxInt = t.FieldOff(2)
		v2.AddArg(dst)
		v3 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v3.Aux = t.FieldType(1)
		v4 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(1).PtrTo())
		v4.AuxInt = t.FieldOff(1)
		v4.AddArg(dst)
		v5 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v5.Aux = t.FieldType(0)
		v6 := b.NewValue0(v.Pos, OpOffPtr, t.FieldType(0).PtrTo())
		v6.AuxInt = 0
		v6.AddArg(dst)
		v5.AddArg3(v6, f0, mem)
		v3.AddArg3(v4, f1, v5)
		v1.AddArg3(v2, f2, v3)
		v.AddArg3(v0, f3, v1)
		return true
	}
	// match: (Store {t} dst (Load src mem) mem)
	// cond: !fe.CanSSA(t.(*types.Type))
	// result: (Move {t} [sizeof(t)] dst src mem)
	for {
		t := v.Aux
		dst := v_0
		if v_1.Op != OpLoad {
			break
		}
		mem := v_1.Args[1]
		src := v_1.Args[0]
		if mem != v_2 || !(!fe.CanSSA(t.(*types.Type))) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = sizeof(t)
		v.Aux = t
		v.AddArg3(dst, src, mem)
		return true
	}
	// match: (Store {t} dst (Load src mem) (VarDef {x} mem))
	// cond: !fe.CanSSA(t.(*types.Type))
	// result: (Move {t} [sizeof(t)] dst src (VarDef {x} mem))
	for {
		t := v.Aux
		dst := v_0
		if v_1.Op != OpLoad {
			break
		}
		mem := v_1.Args[1]
		src := v_1.Args[0]
		if v_2.Op != OpVarDef {
			break
		}
		x := v_2.Aux
		if mem != v_2.Args[0] || !(!fe.CanSSA(t.(*types.Type))) {
			break
		}
		v.reset(OpMove)
		v.AuxInt = sizeof(t)
		v.Aux = t
		v0 := b.NewValue0(v.Pos, OpVarDef, types.TypeMem)
		v0.Aux = x
		v0.AddArg(mem)
		v.AddArg3(dst, src, v0)
		return true
	}
	// match: (Store _ (ArrayMake0) mem)
	// result: mem
	for {
		if v_1.Op != OpArrayMake0 {
			break
		}
		mem := v_2
		v.copyOf(mem)
		return true
	}
	// match: (Store dst (ArrayMake1 e) mem)
	// result: (Store {e.Type} dst e mem)
	for {
		dst := v_0
		if v_1.Op != OpArrayMake1 {
			break
		}
		e := v_1.Args[0]
		mem := v_2
		v.reset(OpStore)
		v.Aux = e.Type
		v.AddArg3(dst, e, mem)
		return true
	}
	// match: (Store (Load (OffPtr [c] (SP)) mem) x mem)
	// cond: isConstZero(x) && mem.Op == OpStaticCall && isSameSym(mem.Aux, "runtime.newobject") && c == config.ctxt.FixedFrameSize() + config.RegSize
	// result: mem
	for {
		if v_0.Op != OpLoad {
			break
		}
		mem := v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpOffPtr {
			break
		}
		c := v_0_0.AuxInt
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpSP {
			break
		}
		x := v_1
		if mem != v_2 || !(isConstZero(x) && mem.Op == OpStaticCall && isSameSym(mem.Aux, "runtime.newobject") && c == config.ctxt.FixedFrameSize()+config.RegSize) {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (Store (OffPtr (Load (OffPtr [c] (SP)) mem)) x mem)
	// cond: isConstZero(x) && mem.Op == OpStaticCall && isSameSym(mem.Aux, "runtime.newobject") && c == config.ctxt.FixedFrameSize() + config.RegSize
	// result: mem
	for {
		if v_0.Op != OpOffPtr {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpLoad {
			break
		}
		mem := v_0_0.Args[1]
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpOffPtr {
			break
		}
		c := v_0_0_0.AuxInt
		v_0_0_0_0 := v_0_0_0.Args[0]
		if v_0_0_0_0.Op != OpSP {
			break
		}
		x := v_1
		if mem != v_2 || !(isConstZero(x) && mem.Op == OpStaticCall && isSameSym(mem.Aux, "runtime.newobject") && c == config.ctxt.FixedFrameSize()+config.RegSize) {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (Store {t1} op1:(OffPtr [o1] p1) d1 m2:(Store {t2} op2:(OffPtr [0] p2) d2 m3:(Move [n] p3 _ mem)))
	// cond: m2.Uses == 1 && m3.Uses == 1 && o1 == sizeof(t2) && n == sizeof(t2) + sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && clobber(m2, m3)
	// result: (Store {t1} op1 d1 (Store {t2} op2 d2 mem))
	for {
		t1 := v.Aux
		op1 := v_0
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v_1
		m2 := v_2
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr || op2.AuxInt != 0 {
			break
		}
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpMove {
			break
		}
		n := m3.AuxInt
		mem := m3.Args[2]
		p3 := m3.Args[0]
		if !(m2.Uses == 1 && m3.Uses == 1 && o1 == sizeof(t2) && n == sizeof(t2)+sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && clobber(m2, m3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v0 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v0.Aux = t2
		v0.AddArg3(op2, d2, mem)
		v.AddArg3(op1, d1, v0)
		return true
	}
	// match: (Store {t1} op1:(OffPtr [o1] p1) d1 m2:(Store {t2} op2:(OffPtr [o2] p2) d2 m3:(Store {t3} op3:(OffPtr [0] p3) d3 m4:(Move [n] p4 _ mem))))
	// cond: m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && o2 == sizeof(t3) && o1-o2 == sizeof(t2) && n == sizeof(t3) + sizeof(t2) + sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && clobber(m2, m3, m4)
	// result: (Store {t1} op1 d1 (Store {t2} op2 d2 (Store {t3} op3 d3 mem)))
	for {
		t1 := v.Aux
		op1 := v_0
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v_1
		m2 := v_2
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpStore {
			break
		}
		t3 := m3.Aux
		_ = m3.Args[2]
		op3 := m3.Args[0]
		if op3.Op != OpOffPtr || op3.AuxInt != 0 {
			break
		}
		p3 := op3.Args[0]
		d3 := m3.Args[1]
		m4 := m3.Args[2]
		if m4.Op != OpMove {
			break
		}
		n := m4.AuxInt
		mem := m4.Args[2]
		p4 := m4.Args[0]
		if !(m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && o2 == sizeof(t3) && o1-o2 == sizeof(t2) && n == sizeof(t3)+sizeof(t2)+sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && clobber(m2, m3, m4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v0 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v0.Aux = t2
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v1.AddArg3(op3, d3, mem)
		v0.AddArg3(op2, d2, v1)
		v.AddArg3(op1, d1, v0)
		return true
	}
	// match: (Store {t1} op1:(OffPtr [o1] p1) d1 m2:(Store {t2} op2:(OffPtr [o2] p2) d2 m3:(Store {t3} op3:(OffPtr [o3] p3) d3 m4:(Store {t4} op4:(OffPtr [0] p4) d4 m5:(Move [n] p5 _ mem)))))
	// cond: m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && m5.Uses == 1 && o3 == sizeof(t4) && o2-o3 == sizeof(t3) && o1-o2 == sizeof(t2) && n == sizeof(t4) + sizeof(t3) + sizeof(t2) + sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && clobber(m2, m3, m4, m5)
	// result: (Store {t1} op1 d1 (Store {t2} op2 d2 (Store {t3} op3 d3 (Store {t4} op4 d4 mem))))
	for {
		t1 := v.Aux
		op1 := v_0
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v_1
		m2 := v_2
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpStore {
			break
		}
		t3 := m3.Aux
		_ = m3.Args[2]
		op3 := m3.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d3 := m3.Args[1]
		m4 := m3.Args[2]
		if m4.Op != OpStore {
			break
		}
		t4 := m4.Aux
		_ = m4.Args[2]
		op4 := m4.Args[0]
		if op4.Op != OpOffPtr || op4.AuxInt != 0 {
			break
		}
		p4 := op4.Args[0]
		d4 := m4.Args[1]
		m5 := m4.Args[2]
		if m5.Op != OpMove {
			break
		}
		n := m5.AuxInt
		mem := m5.Args[2]
		p5 := m5.Args[0]
		if !(m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && m5.Uses == 1 && o3 == sizeof(t4) && o2-o3 == sizeof(t3) && o1-o2 == sizeof(t2) && n == sizeof(t4)+sizeof(t3)+sizeof(t2)+sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && clobber(m2, m3, m4, m5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v0 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v0.Aux = t2
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v2.Aux = t4
		v2.AddArg3(op4, d4, mem)
		v1.AddArg3(op3, d3, v2)
		v0.AddArg3(op2, d2, v1)
		v.AddArg3(op1, d1, v0)
		return true
	}
	// match: (Store {t1} op1:(OffPtr [o1] p1) d1 m2:(Store {t2} op2:(OffPtr [0] p2) d2 m3:(Zero [n] p3 mem)))
	// cond: m2.Uses == 1 && m3.Uses == 1 && o1 == sizeof(t2) && n == sizeof(t2) + sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && clobber(m2, m3)
	// result: (Store {t1} op1 d1 (Store {t2} op2 d2 mem))
	for {
		t1 := v.Aux
		op1 := v_0
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v_1
		m2 := v_2
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr || op2.AuxInt != 0 {
			break
		}
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpZero {
			break
		}
		n := m3.AuxInt
		mem := m3.Args[1]
		p3 := m3.Args[0]
		if !(m2.Uses == 1 && m3.Uses == 1 && o1 == sizeof(t2) && n == sizeof(t2)+sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && clobber(m2, m3)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v0 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v0.Aux = t2
		v0.AddArg3(op2, d2, mem)
		v.AddArg3(op1, d1, v0)
		return true
	}
	// match: (Store {t1} op1:(OffPtr [o1] p1) d1 m2:(Store {t2} op2:(OffPtr [o2] p2) d2 m3:(Store {t3} op3:(OffPtr [0] p3) d3 m4:(Zero [n] p4 mem))))
	// cond: m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && o2 == sizeof(t3) && o1-o2 == sizeof(t2) && n == sizeof(t3) + sizeof(t2) + sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && clobber(m2, m3, m4)
	// result: (Store {t1} op1 d1 (Store {t2} op2 d2 (Store {t3} op3 d3 mem)))
	for {
		t1 := v.Aux
		op1 := v_0
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v_1
		m2 := v_2
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpStore {
			break
		}
		t3 := m3.Aux
		_ = m3.Args[2]
		op3 := m3.Args[0]
		if op3.Op != OpOffPtr || op3.AuxInt != 0 {
			break
		}
		p3 := op3.Args[0]
		d3 := m3.Args[1]
		m4 := m3.Args[2]
		if m4.Op != OpZero {
			break
		}
		n := m4.AuxInt
		mem := m4.Args[1]
		p4 := m4.Args[0]
		if !(m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && o2 == sizeof(t3) && o1-o2 == sizeof(t2) && n == sizeof(t3)+sizeof(t2)+sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && clobber(m2, m3, m4)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v0 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v0.Aux = t2
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v1.AddArg3(op3, d3, mem)
		v0.AddArg3(op2, d2, v1)
		v.AddArg3(op1, d1, v0)
		return true
	}
	// match: (Store {t1} op1:(OffPtr [o1] p1) d1 m2:(Store {t2} op2:(OffPtr [o2] p2) d2 m3:(Store {t3} op3:(OffPtr [o3] p3) d3 m4:(Store {t4} op4:(OffPtr [0] p4) d4 m5:(Zero [n] p5 mem)))))
	// cond: m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && m5.Uses == 1 && o3 == sizeof(t4) && o2-o3 == sizeof(t3) && o1-o2 == sizeof(t2) && n == sizeof(t4) + sizeof(t3) + sizeof(t2) + sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && clobber(m2, m3, m4, m5)
	// result: (Store {t1} op1 d1 (Store {t2} op2 d2 (Store {t3} op3 d3 (Store {t4} op4 d4 mem))))
	for {
		t1 := v.Aux
		op1 := v_0
		if op1.Op != OpOffPtr {
			break
		}
		o1 := op1.AuxInt
		p1 := op1.Args[0]
		d1 := v_1
		m2 := v_2
		if m2.Op != OpStore {
			break
		}
		t2 := m2.Aux
		_ = m2.Args[2]
		op2 := m2.Args[0]
		if op2.Op != OpOffPtr {
			break
		}
		o2 := op2.AuxInt
		p2 := op2.Args[0]
		d2 := m2.Args[1]
		m3 := m2.Args[2]
		if m3.Op != OpStore {
			break
		}
		t3 := m3.Aux
		_ = m3.Args[2]
		op3 := m3.Args[0]
		if op3.Op != OpOffPtr {
			break
		}
		o3 := op3.AuxInt
		p3 := op3.Args[0]
		d3 := m3.Args[1]
		m4 := m3.Args[2]
		if m4.Op != OpStore {
			break
		}
		t4 := m4.Aux
		_ = m4.Args[2]
		op4 := m4.Args[0]
		if op4.Op != OpOffPtr || op4.AuxInt != 0 {
			break
		}
		p4 := op4.Args[0]
		d4 := m4.Args[1]
		m5 := m4.Args[2]
		if m5.Op != OpZero {
			break
		}
		n := m5.AuxInt
		mem := m5.Args[1]
		p5 := m5.Args[0]
		if !(m2.Uses == 1 && m3.Uses == 1 && m4.Uses == 1 && m5.Uses == 1 && o3 == sizeof(t4) && o2-o3 == sizeof(t3) && o1-o2 == sizeof(t2) && n == sizeof(t4)+sizeof(t3)+sizeof(t2)+sizeof(t1) && isSamePtr(p1, p2) && isSamePtr(p2, p3) && isSamePtr(p3, p4) && isSamePtr(p4, p5) && clobber(m2, m3, m4, m5)) {
			break
		}
		v.reset(OpStore)
		v.Aux = t1
		v0 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v0.Aux = t2
		v1 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v1.Aux = t3
		v2 := b.NewValue0(v.Pos, OpStore, types.TypeMem)
		v2.Aux = t4
		v2.AddArg3(op4, d4, mem)
		v1.AddArg3(op3, d3, v2)
		v0.AddArg3(op2, d2, v1)
		v.AddArg3(op1, d1, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpStringLen(v *Value) bool {
	v_0 := v.Args[0]
	// match: (StringLen (StringMake _ (Const64 <t> [c])))
	// result: (Const64 <t> [c])
	for {
		if v_0.Op != OpStringMake {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpConst64 {
			break
		}
		t := v_0_1.Type
		c := v_0_1.AuxInt
		v.reset(OpConst64)
		v.Type = t
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValuegeneric_OpStringPtr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (StringPtr (StringMake (Addr <t> {s} base) _))
	// result: (Addr <t> {s} base)
	for {
		if v_0.Op != OpStringMake {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpAddr {
			break
		}
		t := v_0_0.Type
		s := v_0_0.Aux
		base := v_0_0.Args[0]
		v.reset(OpAddr)
		v.Type = t
		v.Aux = s
		v.AddArg(base)
		return true
	}
	return false
}
func rewriteValuegeneric_OpStructSelect(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	fe := b.Func.fe
	// match: (StructSelect (StructMake1 x))
	// result: x
	for {
		if v_0.Op != OpStructMake1 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (StructSelect [0] (StructMake2 x _))
	// result: x
	for {
		if v.AuxInt != 0 || v_0.Op != OpStructMake2 {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (StructSelect [1] (StructMake2 _ x))
	// result: x
	for {
		if v.AuxInt != 1 || v_0.Op != OpStructMake2 {
			break
		}
		x := v_0.Args[1]
		v.copyOf(x)
		return true
	}
	// match: (StructSelect [0] (StructMake3 x _ _))
	// result: x
	for {
		if v.AuxInt != 0 || v_0.Op != OpStructMake3 {
			break
		}
		_ = v_0.Args[2]
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (StructSelect [1] (StructMake3 _ x _))
	// result: x
	for {
		if v.AuxInt != 1 || v_0.Op != OpStructMake3 {
			break
		}
		_ = v_0.Args[2]
		x := v_0.Args[1]
		v.copyOf(x)
		return true
	}
	// match: (StructSelect [2] (StructMake3 _ _ x))
	// result: x
	for {
		if v.AuxInt != 2 || v_0.Op != OpStructMake3 {
			break
		}
		x := v_0.Args[2]
		v.copyOf(x)
		return true
	}
	// match: (StructSelect [0] (StructMake4 x _ _ _))
	// result: x
	for {
		if v.AuxInt != 0 || v_0.Op != OpStructMake4 {
			break
		}
		_ = v_0.Args[3]
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (StructSelect [1] (StructMake4 _ x _ _))
	// result: x
	for {
		if v.AuxInt != 1 || v_0.Op != OpStructMake4 {
			break
		}
		_ = v_0.Args[3]
		x := v_0.Args[1]
		v.copyOf(x)
		return true
	}
	// match: (StructSelect [2] (StructMake4 _ _ x _))
	// result: x
	for {
		if v.AuxInt != 2 || v_0.Op != OpStructMake4 {
			break
		}
		_ = v_0.Args[3]
		x := v_0.Args[2]
		v.copyOf(x)
		return true
	}
	// match: (StructSelect [3] (StructMake4 _ _ _ x))
	// result: x
	for {
		if v.AuxInt != 3 || v_0.Op != OpStructMake4 {
			break
		}
		x := v_0.Args[3]
		v.copyOf(x)
		return true
	}
	// match: (StructSelect [i] x:(Load <t> ptr mem))
	// cond: !fe.CanSSA(t)
	// result: @x.Block (Load <v.Type> (OffPtr <v.Type.PtrTo()> [t.FieldOff(int(i))] ptr) mem)
	for {
		i := v.AuxInt
		x := v_0
		if x.Op != OpLoad {
			break
		}
		t := x.Type
		mem := x.Args[1]
		ptr := x.Args[0]
		if !(!fe.CanSSA(t)) {
			break
		}
		b = x.Block
		v0 := b.NewValue0(v.Pos, OpLoad, v.Type)
		v.copyOf(v0)
		v1 := b.NewValue0(v.Pos, OpOffPtr, v.Type.PtrTo())
		v1.AuxInt = t.FieldOff(int(i))
		v1.AddArg(ptr)
		v0.AddArg2(v1, mem)
		return true
	}
	// match: (StructSelect [0] (IData x))
	// result: (IData x)
	for {
		if v.AuxInt != 0 || v_0.Op != OpIData {
			break
		}
		x := v_0.Args[0]
		v.reset(OpIData)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Sub16 (Const16 [c]) (Const16 [d]))
	// result: (Const16 [int64(int16(c-d))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst16 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c - d))
		return true
	}
	// match: (Sub16 x (Const16 <t> [c]))
	// cond: x.Op != OpConst16
	// result: (Add16 (Const16 <t> [int64(int16(-c))]) x)
	for {
		x := v_0
		if v_1.Op != OpConst16 {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(-c))
		v.AddArg2(v0, x)
		return true
	}
	// match: (Sub16 <t> (Mul16 x y) (Mul16 x z))
	// result: (Mul16 x (Sub16 <t> y z))
	for {
		t := v.Type
		if v_0.Op != OpMul16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if v_1.Op != OpMul16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				z := v_1_1
				v.reset(OpMul16)
				v0 := b.NewValue0(v.Pos, OpSub16, t)
				v0.AddArg2(y, z)
				v.AddArg2(x, v0)
				return true
			}
		}
		break
	}
	// match: (Sub16 x x)
	// result: (Const16 [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	// match: (Sub16 (Add16 x y) x)
	// result: y
	for {
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if x != v_1 {
				continue
			}
			v.copyOf(y)
			return true
		}
		break
	}
	// match: (Sub16 (Add16 x y) y)
	// result: x
	for {
		if v_0.Op != OpAdd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if y != v_1 {
				continue
			}
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Sub16 x (Sub16 i:(Const16 <t>) z))
	// cond: (z.Op != OpConst16 && x.Op != OpConst16)
	// result: (Sub16 (Add16 <t> x z) i)
	for {
		x := v_0
		if v_1.Op != OpSub16 {
			break
		}
		z := v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpAdd16, t)
		v0.AddArg2(x, z)
		v.AddArg2(v0, i)
		return true
	}
	// match: (Sub16 x (Sub16 z i:(Const16 <t>)))
	// cond: (z.Op != OpConst16 && x.Op != OpConst16)
	// result: (Add16 i (Sub16 <t> x z))
	for {
		x := v_0
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst16 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst16 && x.Op != OpConst16) {
			break
		}
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpSub16, t)
		v0.AddArg2(x, z)
		v.AddArg2(i, v0)
		return true
	}
	// match: (Sub16 (Const16 <t> [c]) (Sub16 x (Const16 <t> [d])))
	// result: (Sub16 (Const16 <t> [int64(int16(c+d))]) x)
	for {
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		if v_1.Op != OpSub16 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst16 || v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpSub16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c + d))
		v.AddArg2(v0, x)
		return true
	}
	// match: (Sub16 (Const16 <t> [c]) (Sub16 (Const16 <t> [d]) x))
	// result: (Add16 (Const16 <t> [int64(int16(c-d))]) x)
	for {
		if v_0.Op != OpConst16 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		if v_1.Op != OpSub16 {
			break
		}
		x := v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst16 || v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		v.reset(OpAdd16)
		v0 := b.NewValue0(v.Pos, OpConst16, t)
		v0.AuxInt = int64(int16(c - d))
		v.AddArg2(v0, x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Sub32 (Const32 [c]) (Const32 [d]))
	// result: (Const32 [int64(int32(c-d))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c - d))
		return true
	}
	// match: (Sub32 x (Const32 <t> [c]))
	// cond: x.Op != OpConst32
	// result: (Add32 (Const32 <t> [int64(int32(-c))]) x)
	for {
		x := v_0
		if v_1.Op != OpConst32 {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(-c))
		v.AddArg2(v0, x)
		return true
	}
	// match: (Sub32 <t> (Mul32 x y) (Mul32 x z))
	// result: (Mul32 x (Sub32 <t> y z))
	for {
		t := v.Type
		if v_0.Op != OpMul32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if v_1.Op != OpMul32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				z := v_1_1
				v.reset(OpMul32)
				v0 := b.NewValue0(v.Pos, OpSub32, t)
				v0.AddArg2(y, z)
				v.AddArg2(x, v0)
				return true
			}
		}
		break
	}
	// match: (Sub32 x x)
	// result: (Const32 [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	// match: (Sub32 (Add32 x y) x)
	// result: y
	for {
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if x != v_1 {
				continue
			}
			v.copyOf(y)
			return true
		}
		break
	}
	// match: (Sub32 (Add32 x y) y)
	// result: x
	for {
		if v_0.Op != OpAdd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if y != v_1 {
				continue
			}
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Sub32 x (Sub32 i:(Const32 <t>) z))
	// cond: (z.Op != OpConst32 && x.Op != OpConst32)
	// result: (Sub32 (Add32 <t> x z) i)
	for {
		x := v_0
		if v_1.Op != OpSub32 {
			break
		}
		z := v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpAdd32, t)
		v0.AddArg2(x, z)
		v.AddArg2(v0, i)
		return true
	}
	// match: (Sub32 x (Sub32 z i:(Const32 <t>)))
	// cond: (z.Op != OpConst32 && x.Op != OpConst32)
	// result: (Add32 i (Sub32 <t> x z))
	for {
		x := v_0
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst32 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst32 && x.Op != OpConst32) {
			break
		}
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpSub32, t)
		v0.AddArg2(x, z)
		v.AddArg2(i, v0)
		return true
	}
	// match: (Sub32 (Const32 <t> [c]) (Sub32 x (Const32 <t> [d])))
	// result: (Sub32 (Const32 <t> [int64(int32(c+d))]) x)
	for {
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		if v_1.Op != OpSub32 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst32 || v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpSub32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c + d))
		v.AddArg2(v0, x)
		return true
	}
	// match: (Sub32 (Const32 <t> [c]) (Sub32 (Const32 <t> [d]) x))
	// result: (Add32 (Const32 <t> [int64(int32(c-d))]) x)
	for {
		if v_0.Op != OpConst32 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		if v_1.Op != OpSub32 {
			break
		}
		x := v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst32 || v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		v.reset(OpAdd32)
		v0 := b.NewValue0(v.Pos, OpConst32, t)
		v0.AuxInt = int64(int32(c - d))
		v.AddArg2(v0, x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Sub32F (Const32F [c]) (Const32F [d]))
	// result: (Const32F [auxFrom32F(auxTo32F(c) - auxTo32F(d))])
	for {
		if v_0.Op != OpConst32F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst32F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst32F)
		v.AuxInt = auxFrom32F(auxTo32F(c) - auxTo32F(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Sub64 (Const64 [c]) (Const64 [d]))
	// result: (Const64 [c-d])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64)
		v.AuxInt = c - d
		return true
	}
	// match: (Sub64 x (Const64 <t> [c]))
	// cond: x.Op != OpConst64
	// result: (Add64 (Const64 <t> [-c]) x)
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = -c
		v.AddArg2(v0, x)
		return true
	}
	// match: (Sub64 <t> (Mul64 x y) (Mul64 x z))
	// result: (Mul64 x (Sub64 <t> y z))
	for {
		t := v.Type
		if v_0.Op != OpMul64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if v_1.Op != OpMul64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				z := v_1_1
				v.reset(OpMul64)
				v0 := b.NewValue0(v.Pos, OpSub64, t)
				v0.AddArg2(y, z)
				v.AddArg2(x, v0)
				return true
			}
		}
		break
	}
	// match: (Sub64 x x)
	// result: (Const64 [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	// match: (Sub64 (Add64 x y) x)
	// result: y
	for {
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if x != v_1 {
				continue
			}
			v.copyOf(y)
			return true
		}
		break
	}
	// match: (Sub64 (Add64 x y) y)
	// result: x
	for {
		if v_0.Op != OpAdd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if y != v_1 {
				continue
			}
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Sub64 x (Sub64 i:(Const64 <t>) z))
	// cond: (z.Op != OpConst64 && x.Op != OpConst64)
	// result: (Sub64 (Add64 <t> x z) i)
	for {
		x := v_0
		if v_1.Op != OpSub64 {
			break
		}
		z := v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpAdd64, t)
		v0.AddArg2(x, z)
		v.AddArg2(v0, i)
		return true
	}
	// match: (Sub64 x (Sub64 z i:(Const64 <t>)))
	// cond: (z.Op != OpConst64 && x.Op != OpConst64)
	// result: (Add64 i (Sub64 <t> x z))
	for {
		x := v_0
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst64 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst64 && x.Op != OpConst64) {
			break
		}
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpSub64, t)
		v0.AddArg2(x, z)
		v.AddArg2(i, v0)
		return true
	}
	// match: (Sub64 (Const64 <t> [c]) (Sub64 x (Const64 <t> [d])))
	// result: (Sub64 (Const64 <t> [c+d]) x)
	for {
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		if v_1.Op != OpSub64 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst64 || v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpSub64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c + d
		v.AddArg2(v0, x)
		return true
	}
	// match: (Sub64 (Const64 <t> [c]) (Sub64 (Const64 <t> [d]) x))
	// result: (Add64 (Const64 <t> [c-d]) x)
	for {
		if v_0.Op != OpConst64 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		if v_1.Op != OpSub64 {
			break
		}
		x := v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst64 || v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		v.reset(OpAdd64)
		v0 := b.NewValue0(v.Pos, OpConst64, t)
		v0.AuxInt = c - d
		v.AddArg2(v0, x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Sub64F (Const64F [c]) (Const64F [d]))
	// result: (Const64F [auxFrom64F(auxTo64F(c) - auxTo64F(d))])
	for {
		if v_0.Op != OpConst64F {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst64F {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst64F)
		v.AuxInt = auxFrom64F(auxTo64F(c) - auxTo64F(d))
		return true
	}
	return false
}
func rewriteValuegeneric_OpSub8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Sub8 (Const8 [c]) (Const8 [d]))
	// result: (Const8 [int64(int8(c-d))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpConst8 {
			break
		}
		d := v_1.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c - d))
		return true
	}
	// match: (Sub8 x (Const8 <t> [c]))
	// cond: x.Op != OpConst8
	// result: (Add8 (Const8 <t> [int64(int8(-c))]) x)
	for {
		x := v_0
		if v_1.Op != OpConst8 {
			break
		}
		t := v_1.Type
		c := v_1.AuxInt
		if !(x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(-c))
		v.AddArg2(v0, x)
		return true
	}
	// match: (Sub8 <t> (Mul8 x y) (Mul8 x z))
	// result: (Mul8 x (Sub8 <t> y z))
	for {
		t := v.Type
		if v_0.Op != OpMul8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if v_1.Op != OpMul8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				z := v_1_1
				v.reset(OpMul8)
				v0 := b.NewValue0(v.Pos, OpSub8, t)
				v0.AddArg2(y, z)
				v.AddArg2(x, v0)
				return true
			}
		}
		break
	}
	// match: (Sub8 x x)
	// result: (Const8 [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	// match: (Sub8 (Add8 x y) x)
	// result: y
	for {
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if x != v_1 {
				continue
			}
			v.copyOf(y)
			return true
		}
		break
	}
	// match: (Sub8 (Add8 x y) y)
	// result: x
	for {
		if v_0.Op != OpAdd8 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			y := v_0_1
			if y != v_1 {
				continue
			}
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Sub8 x (Sub8 i:(Const8 <t>) z))
	// cond: (z.Op != OpConst8 && x.Op != OpConst8)
	// result: (Sub8 (Add8 <t> x z) i)
	for {
		x := v_0
		if v_1.Op != OpSub8 {
			break
		}
		z := v_1.Args[1]
		i := v_1.Args[0]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpAdd8, t)
		v0.AddArg2(x, z)
		v.AddArg2(v0, i)
		return true
	}
	// match: (Sub8 x (Sub8 z i:(Const8 <t>)))
	// cond: (z.Op != OpConst8 && x.Op != OpConst8)
	// result: (Add8 i (Sub8 <t> x z))
	for {
		x := v_0
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		z := v_1.Args[0]
		i := v_1.Args[1]
		if i.Op != OpConst8 {
			break
		}
		t := i.Type
		if !(z.Op != OpConst8 && x.Op != OpConst8) {
			break
		}
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpSub8, t)
		v0.AddArg2(x, z)
		v.AddArg2(i, v0)
		return true
	}
	// match: (Sub8 (Const8 <t> [c]) (Sub8 x (Const8 <t> [d])))
	// result: (Sub8 (Const8 <t> [int64(int8(c+d))]) x)
	for {
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		if v_1.Op != OpSub8 {
			break
		}
		_ = v_1.Args[1]
		x := v_1.Args[0]
		v_1_1 := v_1.Args[1]
		if v_1_1.Op != OpConst8 || v_1_1.Type != t {
			break
		}
		d := v_1_1.AuxInt
		v.reset(OpSub8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c + d))
		v.AddArg2(v0, x)
		return true
	}
	// match: (Sub8 (Const8 <t> [c]) (Sub8 (Const8 <t> [d]) x))
	// result: (Add8 (Const8 <t> [int64(int8(c-d))]) x)
	for {
		if v_0.Op != OpConst8 {
			break
		}
		t := v_0.Type
		c := v_0.AuxInt
		if v_1.Op != OpSub8 {
			break
		}
		x := v_1.Args[1]
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != OpConst8 || v_1_0.Type != t {
			break
		}
		d := v_1_0.AuxInt
		v.reset(OpAdd8)
		v0 := b.NewValue0(v.Pos, OpConst8, t)
		v0.AuxInt = int64(int8(c - d))
		v.AddArg2(v0, x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpTrunc16to8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc16to8 (Const16 [c]))
	// result: (Const8 [int64(int8(c))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c))
		return true
	}
	// match: (Trunc16to8 (ZeroExt8to16 x))
	// result: x
	for {
		if v_0.Op != OpZeroExt8to16 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc16to8 (SignExt8to16 x))
	// result: x
	for {
		if v_0.Op != OpSignExt8to16 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc16to8 (And16 (Const16 [y]) x))
	// cond: y&0xFF == 0xFF
	// result: (Trunc16to8 x)
	for {
		if v_0.Op != OpAnd16 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst16 {
				continue
			}
			y := v_0_0.AuxInt
			x := v_0_1
			if !(y&0xFF == 0xFF) {
				continue
			}
			v.reset(OpTrunc16to8)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpTrunc32to16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc32to16 (Const32 [c]))
	// result: (Const16 [int64(int16(c))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c))
		return true
	}
	// match: (Trunc32to16 (ZeroExt8to32 x))
	// result: (ZeroExt8to16 x)
	for {
		if v_0.Op != OpZeroExt8to32 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpZeroExt8to16)
		v.AddArg(x)
		return true
	}
	// match: (Trunc32to16 (ZeroExt16to32 x))
	// result: x
	for {
		if v_0.Op != OpZeroExt16to32 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc32to16 (SignExt8to32 x))
	// result: (SignExt8to16 x)
	for {
		if v_0.Op != OpSignExt8to32 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpSignExt8to16)
		v.AddArg(x)
		return true
	}
	// match: (Trunc32to16 (SignExt16to32 x))
	// result: x
	for {
		if v_0.Op != OpSignExt16to32 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc32to16 (And32 (Const32 [y]) x))
	// cond: y&0xFFFF == 0xFFFF
	// result: (Trunc32to16 x)
	for {
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst32 {
				continue
			}
			y := v_0_0.AuxInt
			x := v_0_1
			if !(y&0xFFFF == 0xFFFF) {
				continue
			}
			v.reset(OpTrunc32to16)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpTrunc32to8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc32to8 (Const32 [c]))
	// result: (Const8 [int64(int8(c))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c))
		return true
	}
	// match: (Trunc32to8 (ZeroExt8to32 x))
	// result: x
	for {
		if v_0.Op != OpZeroExt8to32 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc32to8 (SignExt8to32 x))
	// result: x
	for {
		if v_0.Op != OpSignExt8to32 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc32to8 (And32 (Const32 [y]) x))
	// cond: y&0xFF == 0xFF
	// result: (Trunc32to8 x)
	for {
		if v_0.Op != OpAnd32 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst32 {
				continue
			}
			y := v_0_0.AuxInt
			x := v_0_1
			if !(y&0xFF == 0xFF) {
				continue
			}
			v.reset(OpTrunc32to8)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpTrunc64to16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc64to16 (Const64 [c]))
	// result: (Const16 [int64(int16(c))])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(int16(c))
		return true
	}
	// match: (Trunc64to16 (ZeroExt8to64 x))
	// result: (ZeroExt8to16 x)
	for {
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpZeroExt8to16)
		v.AddArg(x)
		return true
	}
	// match: (Trunc64to16 (ZeroExt16to64 x))
	// result: x
	for {
		if v_0.Op != OpZeroExt16to64 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc64to16 (SignExt8to64 x))
	// result: (SignExt8to16 x)
	for {
		if v_0.Op != OpSignExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpSignExt8to16)
		v.AddArg(x)
		return true
	}
	// match: (Trunc64to16 (SignExt16to64 x))
	// result: x
	for {
		if v_0.Op != OpSignExt16to64 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc64to16 (And64 (Const64 [y]) x))
	// cond: y&0xFFFF == 0xFFFF
	// result: (Trunc64to16 x)
	for {
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst64 {
				continue
			}
			y := v_0_0.AuxInt
			x := v_0_1
			if !(y&0xFFFF == 0xFFFF) {
				continue
			}
			v.reset(OpTrunc64to16)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpTrunc64to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc64to32 (Const64 [c]))
	// result: (Const32 [int64(int32(c))])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(int32(c))
		return true
	}
	// match: (Trunc64to32 (ZeroExt8to64 x))
	// result: (ZeroExt8to32 x)
	for {
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpZeroExt8to32)
		v.AddArg(x)
		return true
	}
	// match: (Trunc64to32 (ZeroExt16to64 x))
	// result: (ZeroExt16to32 x)
	for {
		if v_0.Op != OpZeroExt16to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpZeroExt16to32)
		v.AddArg(x)
		return true
	}
	// match: (Trunc64to32 (ZeroExt32to64 x))
	// result: x
	for {
		if v_0.Op != OpZeroExt32to64 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc64to32 (SignExt8to64 x))
	// result: (SignExt8to32 x)
	for {
		if v_0.Op != OpSignExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpSignExt8to32)
		v.AddArg(x)
		return true
	}
	// match: (Trunc64to32 (SignExt16to64 x))
	// result: (SignExt16to32 x)
	for {
		if v_0.Op != OpSignExt16to64 {
			break
		}
		x := v_0.Args[0]
		v.reset(OpSignExt16to32)
		v.AddArg(x)
		return true
	}
	// match: (Trunc64to32 (SignExt32to64 x))
	// result: x
	for {
		if v_0.Op != OpSignExt32to64 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc64to32 (And64 (Const64 [y]) x))
	// cond: y&0xFFFFFFFF == 0xFFFFFFFF
	// result: (Trunc64to32 x)
	for {
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst64 {
				continue
			}
			y := v_0_0.AuxInt
			x := v_0_1
			if !(y&0xFFFFFFFF == 0xFFFFFFFF) {
				continue
			}
			v.reset(OpTrunc64to32)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpTrunc64to8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc64to8 (Const64 [c]))
	// result: (Const8 [int64(int8(c))])
	for {
		if v_0.Op != OpConst64 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst8)
		v.AuxInt = int64(int8(c))
		return true
	}
	// match: (Trunc64to8 (ZeroExt8to64 x))
	// result: x
	for {
		if v_0.Op != OpZeroExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc64to8 (SignExt8to64 x))
	// result: x
	for {
		if v_0.Op != OpSignExt8to64 {
			break
		}
		x := v_0.Args[0]
		v.copyOf(x)
		return true
	}
	// match: (Trunc64to8 (And64 (Const64 [y]) x))
	// cond: y&0xFF == 0xFF
	// result: (Trunc64to8 x)
	for {
		if v_0.Op != OpAnd64 {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpConst64 {
				continue
			}
			y := v_0_0.AuxInt
			x := v_0_1
			if !(y&0xFF == 0xFF) {
				continue
			}
			v.reset(OpTrunc64to8)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpXor16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Xor16 (Const16 [c]) (Const16 [d]))
	// result: (Const16 [int64(int16(c^d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst16 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst16)
			v.AuxInt = int64(int16(c ^ d))
			return true
		}
		break
	}
	// match: (Xor16 x x)
	// result: (Const16 [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	// match: (Xor16 (Const16 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Xor16 x (Xor16 x y))
	// result: y
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpXor16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.copyOf(y)
				return true
			}
		}
		break
	}
	// match: (Xor16 (Xor16 i:(Const16 <t>) z) x)
	// cond: (z.Op != OpConst16 && x.Op != OpConst16)
	// result: (Xor16 i (Xor16 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpXor16 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst16 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst16 && x.Op != OpConst16) {
					continue
				}
				v.reset(OpXor16)
				v0 := b.NewValue0(v.Pos, OpXor16, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Xor16 (Const16 <t> [c]) (Xor16 (Const16 <t> [d]) x))
	// result: (Xor16 (Const16 <t> [int64(int16(c^d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst16 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpXor16 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst16 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpXor16)
				v0 := b.NewValue0(v.Pos, OpConst16, t)
				v0.AuxInt = int64(int16(c ^ d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpXor32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Xor32 (Const32 [c]) (Const32 [d]))
	// result: (Const32 [int64(int32(c^d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst32 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst32)
			v.AuxInt = int64(int32(c ^ d))
			return true
		}
		break
	}
	// match: (Xor32 x x)
	// result: (Const32 [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	// match: (Xor32 (Const32 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Xor32 x (Xor32 x y))
	// result: y
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpXor32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.copyOf(y)
				return true
			}
		}
		break
	}
	// match: (Xor32 (Xor32 i:(Const32 <t>) z) x)
	// cond: (z.Op != OpConst32 && x.Op != OpConst32)
	// result: (Xor32 i (Xor32 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpXor32 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst32 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst32 && x.Op != OpConst32) {
					continue
				}
				v.reset(OpXor32)
				v0 := b.NewValue0(v.Pos, OpXor32, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Xor32 (Const32 <t> [c]) (Xor32 (Const32 <t> [d]) x))
	// result: (Xor32 (Const32 <t> [int64(int32(c^d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst32 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpXor32 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst32 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpXor32)
				v0 := b.NewValue0(v.Pos, OpConst32, t)
				v0.AuxInt = int64(int32(c ^ d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpXor64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Xor64 (Const64 [c]) (Const64 [d]))
	// result: (Const64 [c^d])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst64 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst64)
			v.AuxInt = c ^ d
			return true
		}
		break
	}
	// match: (Xor64 x x)
	// result: (Const64 [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConst64)
		v.AuxInt = 0
		return true
	}
	// match: (Xor64 (Const64 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Xor64 x (Xor64 x y))
	// result: y
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpXor64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.copyOf(y)
				return true
			}
		}
		break
	}
	// match: (Xor64 (Xor64 i:(Const64 <t>) z) x)
	// cond: (z.Op != OpConst64 && x.Op != OpConst64)
	// result: (Xor64 i (Xor64 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpXor64 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst64 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst64 && x.Op != OpConst64) {
					continue
				}
				v.reset(OpXor64)
				v0 := b.NewValue0(v.Pos, OpXor64, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Xor64 (Const64 <t> [c]) (Xor64 (Const64 <t> [d]) x))
	// result: (Xor64 (Const64 <t> [c^d]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst64 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpXor64 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst64 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpXor64)
				v0 := b.NewValue0(v.Pos, OpConst64, t)
				v0.AuxInt = c ^ d
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpXor8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Xor8 (Const8 [c]) (Const8 [d]))
	// result: (Const8 [int64(int8(c^d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpConst8 {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpConst8)
			v.AuxInt = int64(int8(c ^ d))
			return true
		}
		break
	}
	// match: (Xor8 x x)
	// result: (Const8 [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	// match: (Xor8 (Const8 [0]) x)
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 || v_0.AuxInt != 0 {
				continue
			}
			x := v_1
			v.copyOf(x)
			return true
		}
		break
	}
	// match: (Xor8 x (Xor8 x y))
	// result: y
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpXor8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if x != v_1_0 {
					continue
				}
				y := v_1_1
				v.copyOf(y)
				return true
			}
		}
		break
	}
	// match: (Xor8 (Xor8 i:(Const8 <t>) z) x)
	// cond: (z.Op != OpConst8 && x.Op != OpConst8)
	// result: (Xor8 i (Xor8 <t> z x))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpXor8 {
				continue
			}
			_ = v_0.Args[1]
			v_0_0 := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_0_0, v_0_1 = _i1+1, v_0_1, v_0_0 {
				i := v_0_0
				if i.Op != OpConst8 {
					continue
				}
				t := i.Type
				z := v_0_1
				x := v_1
				if !(z.Op != OpConst8 && x.Op != OpConst8) {
					continue
				}
				v.reset(OpXor8)
				v0 := b.NewValue0(v.Pos, OpXor8, t)
				v0.AddArg2(z, x)
				v.AddArg2(i, v0)
				return true
			}
		}
		break
	}
	// match: (Xor8 (Const8 <t> [c]) (Xor8 (Const8 <t> [d]) x))
	// result: (Xor8 (Const8 <t> [int64(int8(c^d))]) x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpConst8 {
				continue
			}
			t := v_0.Type
			c := v_0.AuxInt
			if v_1.Op != OpXor8 {
				continue
			}
			_ = v_1.Args[1]
			v_1_0 := v_1.Args[0]
			v_1_1 := v_1.Args[1]
			for _i1 := 0; _i1 <= 1; _i1, v_1_0, v_1_1 = _i1+1, v_1_1, v_1_0 {
				if v_1_0.Op != OpConst8 || v_1_0.Type != t {
					continue
				}
				d := v_1_0.AuxInt
				x := v_1_1
				v.reset(OpXor8)
				v0 := b.NewValue0(v.Pos, OpConst8, t)
				v0.AuxInt = int64(int8(c ^ d))
				v.AddArg2(v0, x)
				return true
			}
		}
		break
	}
	return false
}
func rewriteValuegeneric_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (Zero (Load (OffPtr [c] (SP)) mem) mem)
	// cond: mem.Op == OpStaticCall && isSameSym(mem.Aux, "runtime.newobject") && c == config.ctxt.FixedFrameSize() + config.RegSize
	// result: mem
	for {
		if v_0.Op != OpLoad {
			break
		}
		mem := v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpOffPtr {
			break
		}
		c := v_0_0.AuxInt
		v_0_0_0 := v_0_0.Args[0]
		if v_0_0_0.Op != OpSP || mem != v_1 || !(mem.Op == OpStaticCall && isSameSym(mem.Aux, "runtime.newobject") && c == config.ctxt.FixedFrameSize()+config.RegSize) {
			break
		}
		v.copyOf(mem)
		return true
	}
	// match: (Zero {t1} [n] p1 store:(Store {t2} (OffPtr [o2] p2) _ mem))
	// cond: isSamePtr(p1, p2) && store.Uses == 1 && n >= o2 + sizeof(t2) && clobber(store)
	// result: (Zero {t1} [n] p1 mem)
	for {
		n := v.AuxInt
		t1 := v.Aux
		p1 := v_0
		store := v_1
		if store.Op != OpStore {
			break
		}
		t2 := store.Aux
		mem := store.Args[2]
		store_0 := store.Args[0]
		if store_0.Op != OpOffPtr {
			break
		}
		o2 := store_0.AuxInt
		p2 := store_0.Args[0]
		if !(isSamePtr(p1, p2) && store.Uses == 1 && n >= o2+sizeof(t2) && clobber(store)) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = n
		v.Aux = t1
		v.AddArg2(p1, mem)
		return true
	}
	// match: (Zero {t} [n] dst1 move:(Move {t} [n] dst2 _ mem))
	// cond: move.Uses == 1 && isSamePtr(dst1, dst2) && clobber(move)
	// result: (Zero {t} [n] dst1 mem)
	for {
		n := v.AuxInt
		t := v.Aux
		dst1 := v_0
		move := v_1
		if move.Op != OpMove || move.AuxInt != n || move.Aux != t {
			break
		}
		mem := move.Args[2]
		dst2 := move.Args[0]
		if !(move.Uses == 1 && isSamePtr(dst1, dst2) && clobber(move)) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = n
		v.Aux = t
		v.AddArg2(dst1, mem)
		return true
	}
	// match: (Zero {t} [n] dst1 vardef:(VarDef {x} move:(Move {t} [n] dst2 _ mem)))
	// cond: move.Uses == 1 && vardef.Uses == 1 && isSamePtr(dst1, dst2) && clobber(move, vardef)
	// result: (Zero {t} [n] dst1 (VarDef {x} mem))
	for {
		n := v.AuxInt
		t := v.Aux
		dst1 := v_0
		vardef := v_1
		if vardef.Op != OpVarDef {
			break
		}
		x := vardef.Aux
		move := vardef.Args[0]
		if move.Op != OpMove || move.AuxInt != n || move.Aux != t {
			break
		}
		mem := move.Args[2]
		dst2 := move.Args[0]
		if !(move.Uses == 1 && vardef.Uses == 1 && isSamePtr(dst1, dst2) && clobber(move, vardef)) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = n
		v.Aux = t
		v0 := b.NewValue0(v.Pos, OpVarDef, types.TypeMem)
		v0.Aux = x
		v0.AddArg(mem)
		v.AddArg2(dst1, v0)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt16to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt16to32 (Const16 [c]))
	// result: (Const32 [int64(uint16(c))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(uint16(c))
		return true
	}
	// match: (ZeroExt16to32 (Trunc32to16 x:(Rsh32Ux64 _ (Const64 [s]))))
	// cond: s >= 16
	// result: x
	for {
		if v_0.Op != OpTrunc32to16 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh32Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 16) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt16to64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt16to64 (Const16 [c]))
	// result: (Const64 [int64(uint16(c))])
	for {
		if v_0.Op != OpConst16 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(uint16(c))
		return true
	}
	// match: (ZeroExt16to64 (Trunc64to16 x:(Rsh64Ux64 _ (Const64 [s]))))
	// cond: s >= 48
	// result: x
	for {
		if v_0.Op != OpTrunc64to16 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 48) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt32to64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt32to64 (Const32 [c]))
	// result: (Const64 [int64(uint32(c))])
	for {
		if v_0.Op != OpConst32 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(uint32(c))
		return true
	}
	// match: (ZeroExt32to64 (Trunc64to32 x:(Rsh64Ux64 _ (Const64 [s]))))
	// cond: s >= 32
	// result: x
	for {
		if v_0.Op != OpTrunc64to32 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 32) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt8to16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt8to16 (Const8 [c]))
	// result: (Const16 [int64( uint8(c))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst16)
		v.AuxInt = int64(uint8(c))
		return true
	}
	// match: (ZeroExt8to16 (Trunc16to8 x:(Rsh16Ux64 _ (Const64 [s]))))
	// cond: s >= 8
	// result: x
	for {
		if v_0.Op != OpTrunc16to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh16Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 8) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt8to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt8to32 (Const8 [c]))
	// result: (Const32 [int64( uint8(c))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst32)
		v.AuxInt = int64(uint8(c))
		return true
	}
	// match: (ZeroExt8to32 (Trunc32to8 x:(Rsh32Ux64 _ (Const64 [s]))))
	// cond: s >= 24
	// result: x
	for {
		if v_0.Op != OpTrunc32to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh32Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 24) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValuegeneric_OpZeroExt8to64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt8to64 (Const8 [c]))
	// result: (Const64 [int64( uint8(c))])
	for {
		if v_0.Op != OpConst8 {
			break
		}
		c := v_0.AuxInt
		v.reset(OpConst64)
		v.AuxInt = int64(uint8(c))
		return true
	}
	// match: (ZeroExt8to64 (Trunc64to8 x:(Rsh64Ux64 _ (Const64 [s]))))
	// cond: s >= 56
	// result: x
	for {
		if v_0.Op != OpTrunc64to8 {
			break
		}
		x := v_0.Args[0]
		if x.Op != OpRsh64Ux64 {
			break
		}
		_ = x.Args[1]
		x_1 := x.Args[1]
		if x_1.Op != OpConst64 {
			break
		}
		s := x_1.AuxInt
		if !(s >= 56) {
			break
		}
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteBlockgeneric(b *Block) bool {
	switch b.Kind {
	case BlockIf:
		// match: (If (Not cond) yes no)
		// result: (If cond no yes)
		for b.Controls[0].Op == OpNot {
			v_0 := b.Controls[0]
			cond := v_0.Args[0]
			b.Reset(BlockIf)
			b.AddControl(cond)
			b.swapSuccessors()
			return true
		}
		// match: (If (ConstBool [c]) yes no)
		// cond: c == 1
		// result: (First yes no)
		for b.Controls[0].Op == OpConstBool {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c == 1) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (If (ConstBool [c]) yes no)
		// cond: c == 0
		// result: (First no yes)
		for b.Controls[0].Op == OpConstBool {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c == 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	}
	return false
}
