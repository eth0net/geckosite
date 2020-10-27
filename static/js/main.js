function toggleNav() {
  var style = getComputedStyle(document.getElementById("navbar"));
  if (style.display == "none") {
    document.getElementById("navbar").style.display = "flex";
  } else {
    document.getElementById("navbar").style.display = "";
  }

  document.getElementById("menu-button").classList.toggle("nodisplay");
  document.getElementById("close-button").classList.toggle("nodisplay");
  document.body.classList.toggle("nooverflow");
}
