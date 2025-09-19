function toggleSearchModal() {
    const overlay = document.getElementById("overlay");

    if (overlay.style.display === "flex") {
        overlay.style.display = "none";
    } else {
        overlay.style.display = "flex";
    }
}

const modalForm = document.querySelector(".modal form");
if (modalForm) {
    modalForm.addEventListener("submit", () => {
        document.getElementById("overlay").style.display = "none";
    });
}

document.getElementById("overlay").addEventListener("click", (e) => {
    if (e.target.id === "overlay") {
        e.target.style.display = "none";
    }
});