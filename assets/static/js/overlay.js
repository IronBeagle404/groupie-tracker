function toggleSearchModal() {
    const overlay = document.getElementById('overlay');
    overlay.style.display = overlay.style.display === 'flex' ? 'none' : 'flex';

    const filters = document.querySelector('.filters');
    const modal = overlay.querySelector('.modal');

    if (window.innerWidth <= 768) {
        if (!modal.querySelector('.filters')) {
            const clone = filters.cloneNode(true);
            modal.appendChild(clone);
        }
    }
}

const modalForm = document.querySelector('.modal form');
if (modalForm) {
    modalForm.addEventListener('submit', () => {
        document.getElementById('overlay').style.display = 'none';
    });
}

document.getElementById('overlay').addEventListener('click', (e) => {
    if (e.target.id === 'overlay') overlay.style.display = 'none';
});

window.addEventListener('resize', () => {
    const mainContainer = document.querySelector('.main-container');
    const filters = document.querySelector('.filters');

    if (window.innerWidth > 768 && !mainContainer.contains(filters)) {
        mainContainer.insertBefore(filters, mainContainer.firstChild);
    }
});