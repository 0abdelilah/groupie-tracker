function showDialog(id) {
    document.getElementById(id)?.showModal();
}
function closeDialog(id) {
    document.getElementById(id)?.close();
}
function displayLocations(id) {
    fetch('locations/' + (Number(id) + 1))
        .then(res => res.json())
        .then(data => {
            const content = document.getElementById("LocationsPlace");
            if (data.locations.length > 0) {
                content.textContent = data.locations.join(', ');
            }
        })
        .catch(err => console.error(err));
}

function displayMore(id) {
    fetch("/relation/" + (Number(id) + 1))
        .then(res => res.json())
        .then(data => {
            const content = document.getElementById("RelationPlace");
            if (Object.keys(data.datesLocations).length > 0) {
                const lines = Object.entries(data.datesLocations).map(
                    ([location, dates]) => `${location}: ${dates.join(', ')}`
                );
                content.textContent = lines.join('\n');
            }
        })
        .catch(err => console.error(err));
}

