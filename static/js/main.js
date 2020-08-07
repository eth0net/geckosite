function toggleNav() {
  const nav = document.getElementById("navbar");
  // if (nav.style.height) {
  //   nav.style.height = "";
  // } else {
  //   nav.style.height = "calc(100vh - 69px)";
  // }
  document.getElementById("navbar").classList.toggle("nodisplay");
  document.getElementById("menu-button").classList.toggle("nodisplay");
  document.getElementById("close-button").classList.toggle("nodisplay");
  document.body.classList.toggle("nooverflow");
}
