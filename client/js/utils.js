const refreshTable = async () => { // this function gets all actual values from DB
    
    let tBody = doc.getElementById('table-body');

    let response = await fetch(`${baseUrl}/users`, {
        method: 'GET'
    });

    if (!response.ok){
        message(`Connection error, code ${response.status}, ask sysadmin...`, 'error', 15000);
        return;
    }

    let users = await response.json();
   
    let tRows = users.map(user => {
        return `
            <tr>
                <th scope="row">${user.UserId}</th>
                <td>${user.Name}</td>
                <td>${user.Lastname}</td>
                <td>${countUserAge(user.Birthdate)}</td>
                <td>${user.Birthdate}</td>
                <td>
                    <button class="btn btn-warning" id="user-${user.UserId}" onclick="editItem(${user.UserId})">Edit</button>
                    <button class="btn btn-danger" onclick="deleteItem(${user.UserId})">Delete</button>
                </td>
            </tr>
    `
    })

    tRows = tRows.join('');

    tBody.innerHTML = ''; // reset previous html markup
    tBody.innerHTML = tRows;
}

const dateToIsraeliFormat = (date) => { // date - default formatted in JS
    return date.split('-').reverse().join('/');
}

const israeliDateToJSFormat = (date) => {
    return date.split('/').reverse().join('-');
}

const countUserAge = (date) => {
    // may not work properly due to async
    let d = new Date();
    // day, month, year
    let currentDate = [d.getDate(), d.getMonth() + 1, d.getFullYear()];
    date = date.split('/');
    // day, month, year
    let userDate = [parseInt(date[0]), parseInt(date[1]), parseInt(date[2])];
    let userAge = currentDate[2] - userDate[2];

    // firstly check month, then day - if user already had his birthday
    if (userDate[1] < currentDate[1] && userDate[0] < currentDate[0]){
        userAge -= 1; // if birthday was not yet, age is less by 1!
    }
    

    return userAge; 
}

const message = (text, type = '', showDelay = 3000) => {
    const msgBox = doc.getElementById('msgBox');
    let bgClass;

    if (type === 'error') {
        bgClass = 'alert-danger';
    } else if (type === 'success') {
        bgClass = 'alert-success';
    } else {
        bgClass = 'alert-info'; // neutral blue color
    }

    msgBox.classList.add(bgClass);
    msgBox.innerText = text;
    msgBox.style.display = 'block';

    setTimeout(() => {
        msgBox.style.display = 'none';
        msgBox.classList.remove(bgClass);
        msgBox.innerText = '';
    }, showDelay);

}

const serialize = () => {
    let firstName = doc.getElementById('firstName').value;
    let lastName = doc.getElementById('lastName').value;
    let birthDate = doc.getElementById('birthDate').value;
    birthDate = dateToIsraeliFormat(birthDate);
    if (!firstName || !lastName || !birthDate){
        return false;
    }
    
    let data = {"name": firstName, "lastname": lastName, "birthdate": birthDate};
    data = JSON.stringify(data);
    return data;
}