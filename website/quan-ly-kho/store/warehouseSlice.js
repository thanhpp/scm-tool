import { createSlice } from "@reduxjs/toolkit";


const WarehouseSlice = createSlice({
    name: 'warehouse',
    initialState: { warehouse: [] },
    reducers: {
        saveWarehouse(state, action) {
            state.warehouse = action.payload.data
        }
    }
})

export const warehouseSliceActions = WarehouseSlice.actions
export default WarehouseSlice