"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function (o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
        desc = { enumerable: true, get: function () { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function (o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function (o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function (o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const gojs = __importStar(require("gojs"));
let v = gojs.Uint64(123);
console.log(v);
v = gojs.Uint64(v);
console.log('sum', v.Add(1, 2, 3, gojs.Uint64(4)));
console.log(v.DivMod(2))

console.log('MaxUint64:', gojs.MaxUint64);
console.log('MaxUint32:', gojs.MaxUint32);
console.log('MaxUint16:', gojs.MaxUint16);
console.log('MaxUint8:', gojs.MaxUint8);
console.log('MaxInt64:', gojs.MaxInt64);
console.log('MaxInt32:', gojs.MaxInt32);
console.log('MaxInt16:', gojs.MaxInt16);
console.log('MaxInt8:', gojs.MaxInt8);
console.log('MinInt64:', gojs.MinInt64);
console.log('MinInt32:', gojs.MinInt32);
console.log('MinInt16:', gojs.MinInt16);
console.log('MinInt8:', gojs.MinInt8);
