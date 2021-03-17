let userId_temp = 0;

const addItem = async () => {
    let serializedData = serialize();
    if (!serializedData) {
        message(`Please fill all the fields!`, 'error');
        return;
    }

    let response = await fetch(`${baseUrl}/users`, {
        method: 'POST',
        body: serializedData
    });

    if (response.ok){
        message(`User created successfully!`, 'success');
        refreshTable();
    } else {
        message(`User update error`, 'error');
        return;
    }
}

const deleteItem = async (id) => {
    let response = await fetch(`${baseUrl}/users/${id}`, {
        method: 'DELETE'
    });
    if (response.ok){
        message(`User deleted successfully!`, 'success');
        resetUI();
        refreshTable();
    } else {
        message(`User update error`, 'error');
        return;
    }

}

const editItem = (id) => {
    // Here we just put desired row values to inputs, to edit them
    let editedTableRow = doc.querySelector(`#user-${id}`).parentElement.parentElement;
    let cells = Array.from(editedTableRow.children);
    userId_temp = parseInt(cells[0].innerText);
    let firstName = cells[1].innerText;
    let lastName = cells[2].innerText;
    let birthDate = cells[4].innerText;

    birthDate = israeliDateToJSFormat(birthDate); // for JS, we need to convert date back

    doc.getElementById('firstName').value = firstName;
    doc.getElementById('lastName').value = lastName;
    doc.getElementById('birthDate').value = birthDate;
    // We have 2 options: update or cancel
    btnAdd.disabled = true;
    btnUpdate.disabled = false;
    btnCancel.disabled = false;

}

const updateItem = async () => {
    // Here we just serialize edited input values and send them to the server
    let serializedData = serialize();
    if (!serializedData) {
        message(`Please fill all the fields!`, 'error');
        return;
    }
   
    let response = await fetch(`${baseUrl}/users/${userId_temp}`, {
        method: 'PUT',
        body: serializedData
    });

    if (response.ok){
        message(`User updated successfully!`, 'success');
        resetUI();
        refreshTable();
    } else {
        message(`User update error`, 'error');
        return;
    }
    
}

const resetUI = () => { // reset UI elements after finish updating or cancel updating
    doc.getElementById('firstName').value = '';
    doc.getElementById('lastName').value = '';
    doc.getElementById('birthDate').value = '';
    btnAdd.disabled = false;
    btnUpdate.disabled = true;
    btnCancel.disabled = true;
}
