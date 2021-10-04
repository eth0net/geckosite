"use strict";

const gallery = document.querySelector('#gallery');

const images = [...gallery.children].map(e => ({
    alt: e.getAttribute("alt"),
    src: e.getAttribute("src"),
}))

ReactDOM.render(e(Gallery, {images}), gallery);
