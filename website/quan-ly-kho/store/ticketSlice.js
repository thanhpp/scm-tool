import { createSlice } from "@reduxjs/toolkit";


const TicketSlice = createSlice({
    name: 'ticket',
    initialState: { ticket: [], isSearching: false },
    reducers: {
        ticketSearch(state, action) {
            state.ticket = action.payload.data
        },
        isSearching(state, action) {
            state.isSearching = action.payload.data
        }
    }
})

export const ticketSliceActions = TicketSlice.actions
export default TicketSlice