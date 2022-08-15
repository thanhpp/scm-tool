import { createSlice } from "@reduxjs/toolkit";


const AuthSlice = createSlice({
    name: 'auth',
    initialState: { auth: false },
    reducers: {
        setAuth(state, action) {
            state.auth = action.payload.data
        }
    }
})

export const authSliceActions = AuthSlice.actions
export default AuthSlice