console.log(123)
var t = setTimeout(() => {
    console.log("cb 2")
}, 2000)
setTimeout(() => {
    console.log("cb ok", t)
    clearTimeout(t)

    let v = 0
    setInterval(function () {
        console.log(v++)
        if (v > 4) {
            clearInterval(this)
        }
    }, 1000)
}, 1000)
