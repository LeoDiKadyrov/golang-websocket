import { isValidUsername } from "./lib/validation.js"

const usernameInput = document.getElementsByClassName("recover__login")[0]
const formSubmitButton = document.getElementsByClassName("recover__submit")[0]
const successModal = document.getElementById("successModal")
const successSpan = document.getElementById("closeSuccessModal");
const failureModal = document.getElementById("failureModal")
const failureSpan = document.getElementById("closeFailureModal")

let userInfo = {
  username: "",
  password: ""
}
formSubmitButton.addEventListener("click", async (event) => {
  event.preventDefault()
  userInfo.username = usernameInput.value

  let dataValidated = isValidUsername(userInfo.username)

  if (dataValidated) {
      fetch("recover", {
          method: "POST", 
          headers: {
              "Content-Type": "application/json",
          },
          body: JSON.stringify(userInfo),
      })
          .then((response) => {
              if (!response.ok) {
                failureModal.style.display = "block";
                throw new Error("Password recovery failed")
              }
              return response.text()
          })
          .then((message) => {
              console.log(message);
              successModal.style.display = "block";
          })
          .catch((error) => {
              console.error("Password recovery error: ", error?.message)
          })
  }
  usernameInput.value = ""
})

successSpan.onclick = function() {
    successModal.style.display = "none";
    window.location.replace("/newpassword")
}

failureSpan.onclick = function() {
    failureModal.style.display = "none";
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
    if (event.target == successModal) {
        successModal.style.display = "none";
        window.location.replace("/");
    } else if (event.target == failureModal) {
        failureModal.style.display = "none";
    }
}