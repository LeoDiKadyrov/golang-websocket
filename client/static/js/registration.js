import { createRegAuthSubmitBtnEventListener } from "./lib/formHandler.js";

const usernameInput = document.getElementsByClassName("registration__login")[0];
const passwordInput = document.getElementsByClassName("registration__password")[0];
const formSubmitButton = document.getElementsByClassName("registration__submit")[0];
const modal = document.getElementsByClassName("registration__modal")[0]
const span = document.getElementsByClassName("close")[0];

createRegAuthSubmitBtnEventListener(formSubmitButton, usernameInput, passwordInput, modal, "register");

// When the user clicks on <span> (x), close the modal
span.onclick = function() {
  modal.style.display = "none";
  window.location.replace("/");
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
    if (event.target == modal) {
        modal.style.display = "none";
        window.location.replace("/");
    }
}
