//  **------Tabel 1**
$(function() {
    $('.select-country').select2();
    $('.select-university').select2();
    $('#majorsList').DataTable(); 
});


    /* Formatting function for row details - modify as you need */
function format ( d ) {
    // `d` is the original data object for the row
    return '<table cellpadding="5" cellspacing="0" border="0" style="padding-left:50px;">'+
        '<tr>'+
            '<td>Full name:</td>'+
            '<td>'+d.name+'</td>'+
        '</tr>'+
        '<tr>'+
            '<td>Extension number:</td>'+
            '<td>'+d.extn+'</td>'+
        '</tr>'+
        '<tr>'+
            '<td>Extra info:</td>'+
            '<td>And any further details here (images etc)...</td>'+
        '</tr>'+
    '</table>';
}


// Delete btn js
document.addEventListener('DOMContentLoaded', (event) => {
    // Function to handle delete action
    const handleDelete = (event) => {
        const deleteButton = event.target;
        if (deleteButton.classList.contains('delete-btn')) {
            const row = deleteButton.closest('tr');
            row.remove();
        }
    };
  
    // Add event listener to all delete buttons
    const deleteButtons = document.querySelectorAll('.delete-btn');
    deleteButtons.forEach(button => {
        button.addEventListener('click', handleDelete);
    });
  });