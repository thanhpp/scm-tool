import { createSlice } from "@reduxjs/toolkit";


const WarehouseSlice = createSlice({
    name: 'warehouse',
    initialState: { warehouse: [], isSearching: false },
    reducers: {
        warehouseSearch(state, action) {
            state.warehouse = action.payload.data
        },
        isSearching(state, action) {
            state.isSearching = action.payload.data
        }
    }
})

export const warehouseSliceActions = WarehouseSlice.actions
export default WarehouseSlice