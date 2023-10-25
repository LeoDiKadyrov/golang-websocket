import { inputValidation } from "./lib/validation.js"

const usernameInput = document.getElementsByClassName("registration__login")[0];
const passwordInput = document.getElementsByClassName("registration__password")[0];
const formSubmitButton = document.getElementsByClassName("registration__submit")[0];
const modal = document.getElementsByClassName("registration__modal")[0]
const span = document.getElementsByClassName("close")[0];

let userInfo = {
  username: "",
  password: ""
}
formSubmitButton.addEventListener("click", async (event) => {
  event.preventDefault()
  userInfo.username = usernameInput.value
  userInfo.password = passwordInput.value

  let dataValidated = inputValidation(userInfo.username, userInfo.password)

  if (dataValidated) {
      fetch("authenticate", {
          method: "POST",
          headers: {
              "Content-Type": "application/json",
          },
          body: JSON.stringify(userInfo),
      })
          .then((response) => {
              if (!response.ok) {
                  throw new Error("Registration failed")
              }
              return response.text()
          })
          .then((message) => {
              console.log(message);
              modal.style.display = "block";
          })
          .catch((error) => {
              console.error("Registration error: ", error?.message)
          })
  }
  usernameInput.value = ""
  passwordInput.value = ""
})

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
