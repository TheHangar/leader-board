let year = new Date().getFullYear()
let main = document.querySelector("main")

document.querySelector("#year").innerText = year

main.addEventListener("click", (e) => {
    if (e.target.classList.contains("copy-clipboard")) {
        navigator.clipboard.writeText(e.target.innerText)
    }
})
