"use strict";

const e = React.createElement;

const setVh = () => {
    const vh = window.innerHeight * 0.01;
    document.documentElement.style.setProperty("--vh", `${vh}px`);
};

window.addEventListener("load", setVh);
window.addEventListener("resize", setVh);

function toggleNav() {
    let style = getComputedStyle(document.getElementById("navbar"));
    if (style.display === "none") {
        document.getElementById("navbar").style.display = "flex";
    } else {
        document.getElementById("navbar").style.display = "";
    }

    document.getElementById("menu-button").classList.toggle("nodisplay");
    document.getElementById("close-button").classList.toggle("nodisplay");
    document.body.classList.toggle("nooverflow");
}
