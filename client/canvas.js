window.addEventListener("load", () => {
    const canvas = document.querySelector("#canvas")
    const ctx = canvas.getContext("2d")

    canvas.height = window.innerHeight / 2
    canvas.width = window.innerWidth / 2
})