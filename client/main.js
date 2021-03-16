const doc = document;
const btnAdd = doc.getElementById('btnAdd');
const btnUpdate = doc.getElementById('btnUpdate');
const btnCancel = doc.getElementById('btnCancel');

btnAdd.addEventListener('click', addItem);
btnUpdate.addEventListener('click', updateItem);
btnCancel.addEventListener('click', finishUpdating);