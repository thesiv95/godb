const doc = document;
const baseUrl = 'http://localhost:8000';

const btnAdd = doc.getElementById('btnAdd');
const btnUpdate = doc.getElementById('btnUpdate');
const btnCancel = doc.getElementById('btnCancel');

btnAdd.addEventListener('click', addItem);
btnUpdate.addEventListener('click', updateItem);
btnCancel.addEventListener('click', resetUI);

doc.addEventListener('DOMContentLoaded', refreshTable);