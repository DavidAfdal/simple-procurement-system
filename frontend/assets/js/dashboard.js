$(document).ready(function() {
    loadInventory();

    function loadInventory() {
        apiRequest('/items', 'GET', {}, function(res) {
            const tbody = $('#inventoryTableBody');
            tbody.empty();
           res.data.forEach((item, index) => {
                tbody.append(`
                    <tr>
                        <td class="border px-4 py-2 text-center">${index + 1}</td>
                        <td class="border px-4 py-2">${item.name}</td>
                        <td class="border px-4 py-2 text-center">${item.stock}</td>
                    </tr>
                `);
            });
            
        });
    }
});