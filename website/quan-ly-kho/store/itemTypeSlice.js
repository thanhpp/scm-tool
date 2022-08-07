import { createSlice } from "@reduxjs/toolkit";


const ItemTypeSlice = createSlice({
    name: 'itemtype',
    initialState: { itemType: [], isSearching: false },
    reducers: {
        itemTypeSearch(state, action) {
            state.itemType = action.payload.data
        },
        isSearching(state, action) {
            state.isSearching = action.payload.data
        }
    }
})

export const itemTypeSliceActions = ItemTypeSlice.actions
export default ItemTypeSlice