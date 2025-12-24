
let items = [];
let cart = [];


$(document).ready(function() {
    loadSuppliers();
    loadItems();

    $('#submitOrderBtn').on('click', submitOrder);
    $('#itemsList').on('click', '.addToCartBtn', addToCart);
    $('#cartItems').on('click', '.qty-increase', increaseQty);
    $('#cartItems').on('click', '.qty-decrease', decreaseQty);
    $('#cartItems').on('click', '.removeItemBtn', removeFromCart);

    function renderItems() {
        const container = $('#itemsList');
        container.empty();

        if (items.length === 0) {
            container.html(`
                        <div class="text-center py-16">
                            <div class="bg-blue-50 rounded-full w-20 h-20 flex items-center justify-center mx-auto mb-4">
                                <i class="fas fa-box-open text-blue-300 text-4xl"></i>
                            </div>
                            <p class="text-gray-500 font-medium">No items available</p>
                        </div>
                    `);
            return;
        }

        items.forEach(item => {
            const inCart = cart.find(c => c.itemId == item.id);
            container.append(`
                        <div class="item-card rounded-xl p-5 mb-4 shadow-sm">
                            <div class="flex justify-between items-center">
                                <div class="flex-1">
                                    <h3 class="font-bold text-gray-900 text-lg mb-2">${item.name}</h3>
                                    <div class="flex items-center space-x-3">
                                        <span class="badge-modern bg-blue-100 text-blue-700">
                                            <i class="fas fa-boxes text-xs mr-1"></i>
                                            Stock: ${item.stock}
                                        </span>
                                        <span class="badge-modern bg-green-100 text-green-700">
                                            <i class="fas fa-tag text-xs mr-1"></i>
                                            Rp ${Number(item.price).toLocaleString('id-ID')}
                                        </span>
                                        ${inCart ? `<span class="badge-modern bg-purple-100 text-purple-700">
                                            <i class="fas fa-check-circle text-xs mr-1"></i>
                                            In Cart
                                        </span>` : ''}
                                    </div>
                                </div>
                                <button class="addToCartBtn btn-blue text-white px-5 py-3 rounded-xl font-semibold flex items-center space-x-2" data-id="${item.id}" data-name="${item.name}" data-stock="${item.stock}" data-price="${item.price}">
                                    <i class="fas fa-plus"></i>
                                    <span>Add</span>
                                </button>
                            </div>
                        </div>
                    `);
        });
    }


    function loadSuppliers() {
        apiRequest('/suppliers', 'GET', {}, function(res){
            const select = $('#supplierSelect');
            select.empty().append('<option value="">Choose a supplier...</option>');
            res.data.forEach(s => select.append(`<option value="${s.id}">${s.name}</option>`));
        });
    }

    function loadItems() {
        apiRequest('/items', 'GET', {}, function(res){
            items = res.data;
            console.log(items)
            renderItems();
            updateStats();
        });
    }


    function submitOrder() {
        if (cart.length === 0) {
                    Swal.fire({
                        icon: 'info',
                        title: 'Cart is empty!',
                        text: 'Please add items to cart first',
                        confirmButtonColor: '#3b82f6'
                    });
                    return;
        }

        const supplierId = $('#supplierSelect').val();

                if (!supplierId) {
                    Swal.fire({
                        icon: 'error',
                        title: 'Supplier required!',
                        text: 'Please select a supplier first',
                        confirmButtonColor: '#3b82f6'
                    });
                    return;
                }

                const purchaseDate = $('#purchaseDate').val();

                if (!purchaseDate) {
                    Swal.fire({
                        icon: 'error',
                        title: 'Purchase Date required!',
                        text: 'Please select a purcahse date first',
                        confirmButtonColor: '#3b82f6'
                    });
                    return;
                }

                const payload = {
                    date: purchaseDate,
                    supplier_id: supplierId,
                    items: cart.map(item => ({
                        item_id: item.itemId,
                        qty: item.qty
                    }))
                };

                apiRequest('/purchasings', 'POST', payload, function(res){
                    Swal.fire({
                        icon: 'success',
                        title: 'Order submitted!',
                        text: 'Your purchase order has been created successfully',
                        confirmButtonColor: '#3b82f6'
                    });
                    cart = [];
                    renderCart();
                    renderItems();
                    updateStats();
                    $('#supplierSelect').val('');
                    $('#purchaseDate').val('');
                    loadItems();
                });
    }

    function addToCart() {
            const itemId = $(this).data('id');
            const itemName = $(this).data('name');
            const stock = $(this).data('stock');
            const price = $(this).data('price');

            const existingItem = cart.find(c => c.itemId == itemId);

            if (existingItem) {
                existingItem.qty++;
            } else {
                cart.push({ itemId, itemName, price, stock, qty: 1 });
            }

            renderCart();
            renderItems();
            updateStats();
                
            Swal.fire({
                    icon: 'success',
                    title: 'Added to cart!',
                    text: itemName,
                    timer: 1500,
                    showConfirmButton: false,
                    toast: true,
                    position: 'top-end',
                    background: '#3b82f6',
                    color: '#fff'
            });
    }

    function renderCart() {
        const container = $('#cartItems');

        if (cart.length === 0) {
            container.html(`
                        <div class="text-center py-16">
                            <div class="bg-blue-50 rounded-full w-20 h-20 flex items-center justify-center mx-auto mb-4">
                                <i class="fas fa-shopping-cart text-blue-300 text-4xl"></i>
                            </div>
                            <p class="text-gray-500 font-medium">Your cart is empty</p>
                            <p class="text-gray-400 text-xs mt-2">Start adding items!</p>
                        </div>
                    `);

            $('#cartSummary').hide();
            $('#grandTotal').text('Rp 0');
            return;
        }

        container.empty();
        let grandTotal = 0;
        cart.forEach((item, idx) => {
            const subtotal = item.price * item.qty;
            grandTotal += subtotal;
            container.append(`
                        <div class="bg-gradient-to-br from-blue-50 to-cyan-50 rounded-xl p-4 mb-3 border-2 border-blue-100 shadow-sm">
                            <div class="flex justify-between items-start mb-3">
                                <div class="flex-1">
                                    <h4 class="font-bold text-gray-900 text-sm mb-1">${item.itemName}</h4>
                                    <p class="text-xs text-gray-600 font-medium">Stock: ${item.stock}</p>
                                    <p class="text-sm text-blue-600 font-bold mt-1">Rp ${Number(item.price).toLocaleString('id-ID')} / unit</p>
                                </div>
                                <button class="removeItemBtn text-red-500 hover:text-red-700 hover:bg-red-50 rounded-lg p-2 transition" data-idx="${idx}">
                                    <i class="fas fa-trash-alt"></i>
                                </button>
                            </div>
                            <div class="flex items-center justify-between bg-white rounded-xl p-2 border-2 border-blue-100 shadow-sm mb-3">
                                <button class="qty-decrease qty-btn bg-gray-100 hover:bg-gray-200 text-gray-700 w-9 h-9 rounded-lg flex items-center justify-center font-bold shadow-sm" data-idx="${idx}">
                                    <i class="fas fa-minus text-xs"></i>
                                </button>
                                <span class="font-bold text-gray-900 text-lg mx-4">${item.qty}</span>
                                <button class="qty-increase qty-btn btn-blue text-white w-9 h-9 rounded-lg flex items-center justify-center font-bold" data-idx="${idx}">
                                    <i class="fas fa-plus text-xs"></i>
                                </button>
                            </div>
                            <div class="bg-white rounded-lg p-2 border border-blue-200">
                                <div class="flex justify-between items-center">
                                    <span class="text-xs text-gray-600 font-semibold">Subtotal:</span>
                                    <span class="text-sm font-bold text-blue-600">Rp ${Number(subtotal).toLocaleString('id-ID')}</span>
                                </div>
                            </div>
                        </div>
                    `);
        });

        $('#cartSummary').show();
        $('#grandTotal').text(`Rp ${Number(grandTotal).toLocaleString('id-ID')}`);
        updateGrandTotal();
    }

    function increaseQty() {
        const idx = $(this).data('idx');
        cart[idx].qty++;
        renderCart();
        updateStats();
    }

    function decreaseQty() {
        const idx = $(this).data('idx');
        if (cart[idx].qty > 1) {
            cart[idx].qty--;
            renderCart();
            updateStats();
        }
    }

    function removeFromCart() {
        const idx = $(this).data('idx');
        cart.splice(idx, 1);
        renderCart();
        renderItems();
        updateStats();
    }

    function updateStats() {
        const totalQty = cart.reduce((sum, item) => sum + item.qty, 0);
        $('#totalItems').text(items.length);
        $('#cartCount').text(cart.length);
        $('#totalQty').text(totalQty);
        $('#cartBadge').text(`${cart.length} item${cart.length !== 1 ? 's' : ''}`);
    }

    function updateGrandTotal() {
        const grandTotal = cart.reduce((sum, item) => sum + (item.price * item.qty), 0);
        $('#grandTotal').text(`Rp ${Number(grandTotal).toLocaleString('id-ID')}`);
    }
});