declare module "gojs" {
    export type NumberLike = number | string | Uint64 | Uint32 | Uint16 | Uint8 | Int64 | Int32 | Int16 | Int8
    export interface Number<T> {
        ToNumber(): number

        ToUint64(): Uint64
        ToUint32(): Uint32
        ToUint16(): Uint16
        ToUint8(): Uint8
        ToInt64(): Int64
        ToInt32(): Int32
        ToInt16(): Int16
        ToInt8(): Int8

        Add(...vals: Array<NumberLike>): T
        Sub(...vals: Array<NumberLike>): T
        Mul(...vals: Array<NumberLike>): T
        Div(...vals: Array<NumberLike>): T
        Mod(...vals: Array<NumberLike>): T
        DivMod(val: NumberLike): [T, T]

        /**
         * 
         * if self == o return 0
         * elif self < o return -1
         * else return 1
         */
        Cmp(o: NumberLike): number

        /**
         * reutrn -self
         */
        Neg(): T
        /**
         * reutrn ^self
         */
        Not(): T
        /**
         * reutrn self | ...vals
         */
        Or(...vals: Array<NumberLike>): T
        /**
         * reutrn self ^ ...vals
         */
        Xor(...vals: Array<NumberLike>): T
        /**
         * reutrn self & ...vals
         */
        And(...vals: Array<NumberLike>): T
    }
    export interface Uint64 extends Number<Uint64> {
    }
    export interface Uint32 extends Number<Uint32> {
    }
    export interface Uint16 extends Number<Uint16> {
    }
    export interface Uint8 extends Number<Uint8> {
    }
    export interface Int64 extends Number<Int64> {
    }
    export interface Int32 extends Number<Int32> {
    }
    export interface Int16 extends Number<Int16> {
    }
    export interface Int8 extends Number<Int8> {
    }
    export const MaxUint64: Uint64
    export const MaxUint32: Uint32
    export const MaxUint16: Uint16
    export const MaxUint8: Uint8
    export const MaxInt64: Int64
    export const MaxInt32: Int32
    export const MaxInt16: Int16
    export const MaxInt8: Int8
    export const MinInt64: Int64
    export const MinInt32: Int32
    export const MinInt16: Int16
    export const MinInt8: Int8

    export function Uint64(val?: NumberLike): Uint64
    export function Uint32(val?: NumberLike): Uint32
    export function Uint16(val?: NumberLike): Uint16
    export function Uint8(val?: NumberLike): Uint8
    export function Int64(val?: NumberLike): Int64
    export function Int32(val?: NumberLike): Int32
    export function Int16(val?: NumberLike): Int16
    export function Int8(val?: NumberLike): Int8
}