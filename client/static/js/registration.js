const usernameInput = document.getElementsByClassName("registration__login")[0];
const passwordInput = document.getElementsByClassName("registration__password")[0];
const formSubmitButton = document.getElementsByClassName("registration__submit")[0];
const modal = document.getElementsByClassName("registration__modal")[0]
const span = document.getElementsByClassName("close")[0];

const UserInput = {
    username: "",
    password: ""
}

formSubmitButton.addEventListener("click", async (event) => {
    event.preventDefault();
    UserInput.username = usernameInput.value;
    UserInput.password = passwordInput.value;

    let dataValidated = inputValidation(UserInput.username, UserInput.password)

    if (dataValidated) {
        fetch("/register", { // TODO: Is there a way to handle it better rather than having /register and /registration ???
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(UserInput),
        })
            .then((response) => {
                if (!response.ok) {
                    throw new Error("Registration failed")
                }
                return response.text();
            })
            .then((message) => {
                console.log(message);
                modal.style.display = "block";
            })
            .catch((error) => {
                console.error("Registration error: ", error?.message);
            })

            usernameInput.value = "";
            passwordInput.value = "";
            passwordInput.blur() // unfocus input
            usernameInput.blur() // unfocus input
    } else {
        alert("There is an error during registration")
    }

})

function inputValidation(username, password) {
    if (!isValidUsername(username)) {
        alert("Invalid username. Use only English letters.");
        return false;
    }

    if (!isValidPassword(password)) {
        alert("Invalid password. Password must be at least 8 characters long, contain at least one uppercase letter, one number, and one special character.");
        return false;
    }

    return true; 
}

function isValidUsername(username) {
    if (username.length <= 4) {
        return false;
    }
    const letters = /^[a-zA-Z]+$/;
    return letters.test(username);
}

function isValidPassword(password) {
    if (password.length < 8 || password.length > 40) {
        return false;
    }

    let hasUpperCase = false;
    for (const char of password) {
        if (char >= 'A' && char <= 'Z') {
            hasUpperCase = true;
            break;
        }
    }
    if (!hasUpperCase) {
        return false;
    }

    let hasNumber = false;
    for (const char of password) {
        if (char >= '0' && char <= '9') {
            hasNumber = true;
            break;
        }
    }
    if (!hasNumber) {
        return false;
    }

    const specialCharacters = "!@#$%^&*()_-+=<>?/[]{}|";
    let hasSpecialCharacter = false;
    for (const char of password) {
        if (specialCharacters.includes(char)) {
            hasSpecialCharacter = true;
            break;
        }
    }
    if (!hasSpecialCharacter) {
        return false;
    }

    return true;
}

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
