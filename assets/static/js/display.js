document.addEventListener("DOMContentLoaded", () => {
    const grid = document.querySelector(".artists-grid");
    const cards = grid.querySelectorAll(".artist-card");

    if (cards.length < 3) {
        grid.classList.add("few-cards");
    }
});