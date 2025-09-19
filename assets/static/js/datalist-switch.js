
/* Detect usage of firefox and switch the suggestions datalist format */
document.addEventListener("DOMContentLoaded", () => {
    if (!navigator.userAgent.toLowerCase().includes("firefox")) return;

    const datalist = document.getElementById("suggestions-dl");
    if (!datalist) return;

    const options = Array.from(datalist.querySelectorAll("option"));
    options.forEach(opt => {
        opt.textContent = `${opt.value} - ${opt.getAttribute("data-type")}`;
    });
});
