import { createSlice } from "@reduxjs/toolkit";


const UserSlice = createSlice({
    name: 'user',
    initialState: { user: [], isSearching: false },
    reducers: {
        userSearch(state, action) {
            state.user = action.payload.data
        },
        isSearching(state, action) {
            state.isSearching = action.payload.data
        }
    }
})

export const userSliceActions = UserSlice.actions
export default UserSlice