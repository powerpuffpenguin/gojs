function main() {
    let v = uint64()
    console.log(v)
    println(v)
}
console.log(123)
var t = setTimeout(() => {
    console.log("cb 2")
}, 200)
setTimeout(() => {
    console.log("cb ok", t)
    clearTimeout(t)

    let v = 0
    setInterval(function () {
        console.log(v++)
        if (v > 4) {
            clearInterval(this)
            main()
        }
    }, 100)
}, 100)
