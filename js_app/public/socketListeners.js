socket.on('availableOffers', offers => {
    console.log(offers)
    createOfferEls(offers)
})

socket.on('newOfferAwaiting', offers => {
    console.log(offers) // Added for debugging
    createOfferEls(offers) // Fixed: using the correct variable name
})

socket.on('answerResponse', offerObj => {
    console.log(offerObj)
    addAnswer(offerObj)
})

socket.on('receivedIceCandidateFromServer', iceCandidate => {
    addNewIceCandidate(iceCandidate)
    console.log(iceCandidate)
})

function createOfferEls(offers) { // Fixed: parameter name matches usage
    const answerEl = document.querySelector('#answer')
    
    // Clear existing offers to prevent duplicates
    answerEl.innerHTML = ''
    
    offers.forEach(o => {
        console.log(o)
        const newOfferEl = document.createElement('div')
        // Fixed: closed the class attribute properly
        newOfferEl.innerHTML = `<button class='btn btn-success col-1'> Answer ${o.offererUserName}</button>`
        newOfferEl.addEventListener('click', () => answerOffer(o))
        answerEl.appendChild(newOfferEl)
    })
}
