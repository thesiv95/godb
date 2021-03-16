const addItem = () => {
    let serialized = serialize();
    if (!serialized) {
        message(`Please fill all the fields!`, 'error');
        return;
    }
    message(`New user added successfully!`, 'success');
}

const deleteItem = (id) => {
    message(`User deleted successfully!`, 'success');
}

const editItem = (id) => {
    console.log('editItem', id);
    let editedTableRow = doc.querySelector(`#user-${id}`).parentElement.parentElement;
    let cells = Array.from(editedTableRow.children);
    let firstName = cells[1].innerText;
    let lastName = cells[2].innerText;
    let birthDate = cells[4].innerText;

    birthDate = israeliDateToJSFormat(birthDate); // for JS, we need to convert date back

    doc.getElementById('firstName').value = firstName;
    doc.getElementById('lastName').value = lastName;
    doc.getElementById('birthDate').value = birthDate;

    btnAdd.disabled = true;
    btnUpdate.disabled = false;
    btnCancel.disabled = false;
}

const updateItem = () => {
    let serialized = serialize();
    if (!serialized) {
        message(`Please fill all the fields!`, 'error');
        return;
    }
    message(`User updated successfully!`, 'success');
    finishUpdating();
}

const finishUpdating = () => { // finish or cancel
    doc.getElementById('firstName').value = '';
    doc.getElementById('lastName').value = '';
    doc.getElementById('birthDate').value = '';
    btnAdd.disabled = false;
    btnUpdate.disabled = true;
    btnCancel.disabled = true;
}
