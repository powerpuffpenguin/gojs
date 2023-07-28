import * as gojs from "gojs";

let v = gojs.Uint64(123)
console.log(v)
v = gojs.Uint64(v)
console.log(v)

console.log('MaxUint64:', gojs.MaxUint64)
console.log('MaxUint32:', gojs.MaxUint32)
console.log('MaxUint16:', gojs.MaxUint16)
console.log('MaxUint8:', gojs.MaxUint8)
console.log('MaxInt64:', gojs.MaxInt64)
console.log('MaxInt32:', gojs.MaxInt32)
console.log('MaxInt16:', gojs.MaxInt16)
console.log('MaxInt8:', gojs.MaxInt8)
console.log('MinInt64:', gojs.MinInt64)
console.log('MinInt32:', gojs.MinInt32)
console.log('MinInt16:', gojs.MinInt16)
console.log('MinInt8:', gojs.MinInt8)