// Fonction pour basculer entre le mode sombre et normal
function toggleTheme() {
    const checkbox = document.getElementById('theme-toggle');
    if (checkbox.checked) {
        document.body.classList.add('dark-mode');
    } else {
        document.body.classList.remove('dark-mode');
    }
}

// Fonction pour afficher le tableau des utilisateurs lorsque "Option 1" est cliqué
document.getElementById("Option1").addEventListener("click", function() {
    const userTableContainer = document.getElementById("user-table-container");
    
    // Charger le fichier HTML pour la gestion des utilisateurs
    fetch('/assets/templates/user.html')
        .then(response => response.text())
        .then(data => {
            userTableContainer.innerHTML = data;  // Injecter le contenu de user-management.html dans le conteneur
            userTableContainer.style.display = 'block';  // Afficher le tableau
        })
        .catch(error => console.error('Error loading user-management.html:', error));
});
// Afficher le formulaire pour ajouter un utilisateur
// Afficher le formulaire pour ajouter un utilisateur
function showAddUserForm() {
    document.getElementById("add-user-form").style.display = 'block';  // Affiche le formulaire
}


// Fonction pour vérifier si le username existe déjà dans le tableau
function isUsernameTaken(username) {
    const rows = document.getElementById('user-table').getElementsByTagName('tbody')[0].rows;
    for (let i = 0; i < rows.length; i++) {
        const existingUsername = rows[i].cells[1].textContent; // Récupère le username de la colonne
        if (existingUsername === username) {
            return true;  // Si le username existe déjà dans le tableau
        }
    }
    return false;  // Si le username n'existe pas
}

// Fonction pour ajouter un utilisateur
async function addUser(event) {
    event.preventDefault(); // Empêcher l'envoi du formulaire pour une soumission standard

    // Récupérer les valeurs du formulaire
    const fullName = document.getElementById('full-name').value;
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const role = document.getElementById('role').value;

    // Vérifier si le username est déjà dans le tableau
    if (isUsernameTaken(username)) {
        alert('Username is already taken in the table. Please choose another one.');
        return;  // Ne pas continuer si le username est déjà pris dans le tableau
    }

    // Ajouter l'utilisateur au tableau HTML
    const userTableBody = document.getElementById('user-table').getElementsByTagName('tbody')[0];
    const newRow = userTableBody.insertRow();

    newRow.innerHTML = `
        <td>${fullName}</td>
        <td>${username}</td>
        <td>${role}</td>
        <td><button onclick="deleteUser(this)">Delete</button></td>
    `;

    // Envoi des données au backend pour ajouter l'utilisateur
    fetch('/add-user', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            fullName: fullName,
            username: username,
            password: password,
            role: role
        })
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            console.log('User added to database');
        } else {
            console.error('Failed to add user to database');
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });

    // Réinitialiser le formulaire
    document.querySelector('form').reset();

    // Masquer le formulaire après l'ajout
    document.getElementById('add-user-form').style.display = 'none';
}


// Fonction pour annuler l'ajout d'utilisateur
function cancelAddUser() {
    document.getElementById('add-user-form').style.display = 'none';
}

// Fonction pour supprimer un utilisateur
function deleteUser(button) {
    const row = button.parentElement.parentElement;
    row.remove();
    // Optionnel : supprimer également de la base de données via une autre requête fetch
}

// Fonction pour charger les utilisateurs existants dans le tableau
function loadExistingUsers() {
    fetch('/get-users') // Appel à la route pour récupérer les utilisateurs
        .then(response => response.json())
        .then(users => {
            const userTableBody = document.getElementById('user-table').getElementsByTagName('tbody')[0];
            userTableBody.innerHTML = ''; // Vider le tableau avant de remplir

            // Ajouter les utilisateurs au tableau
            users.forEach(user => {
                const newRow = userTableBody.insertRow();
                newRow.innerHTML = `
                    <td>${user.username}</td>
                    <td>${user.role}</td>
                    <td><button onclick="deleteUser(this)">Delete</button></td>
                `;
            });
        })
        .catch(error => {
            console.error('Error loading users:', error);
        });
}

// Fonction pour afficher le formulaire d'ajout d'utilisateur
function showAddUserForm() {
    document.getElementById("add-user-form").style.display = 'block'; // Affiche le formulaire
}

// Fonction pour annuler l'ajout d'utilisateur
function cancelAddUser() {
    document.getElementById("add-user-form").style.display = 'none'; // Cache le formulaire
}