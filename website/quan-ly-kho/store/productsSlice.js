import { createSlice } from "@reduxjs/toolkit";


const ProductsSice = createSlice({
    name: 'products',
    initialState: { products: [] },
    reducers: {
        saveProducts(state, action) {
            state.products = action.payload.data
        },
    }
})

export const productsSliceActions = ProductsSice.actions
export default ProductsSice