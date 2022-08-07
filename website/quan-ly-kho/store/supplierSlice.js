import { createSlice } from "@reduxjs/toolkit";


const SupplierSlice = createSlice({
    name: 'supplier',
    initialState: { supplier: [], isSearching: false },
    reducers: {
        supplierSearch(state, action) {
            state.supplier = action.payload.data
        },
        isSearching(state, action) {
            state.isSearching = action.payload.data
        }
    }
})

export const supplierSliceActions = SupplierSlice.actions
export default SupplierSlice