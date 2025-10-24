function closePopup(event, id) {
    // Close if clicked on overlay
    if (event.target.classList.contains('popup-overlay')) {
        document.getElementById(id).checked = false;
    }
}

// Attach escape key listener once
document.addEventListener('keydown', (e) => {
    if (e.key.toLowerCase() === 'escape') {
        // Close all popups (or target specific ones if you store the active id)
        document.querySelectorAll('.popup-toggle:checked').forEach((checkbox) => {
            checkbox.checked = false;
        });
    }
});