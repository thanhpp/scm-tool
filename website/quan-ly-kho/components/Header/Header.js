import React from 'react'
import { Avatar } from '@mui/material'
import Link from 'next/link'

function Header() {
    return (
        <div className=' flex justify-between px-[30px] py-[30px] text-[20px] border-b border-white'>
            <div><Link href='/'><span className=' cursor-pointer hover:opacity-80'>Trang chủ</span></Link></div>
            <div><Link href='/warehouse-management'><span className=' cursor-pointer hover:opacity-80'>Quản lý kho</span></Link></div>
            <div><Link href='/products-management'><span className=' cursor-pointer hover:opacity-80'>Quản lý sản phẩm</span></Link></div>
            <div><Link href='/io-management'><span className=' cursor-pointer hover:opacity-80'>Quản lý nhập xuất</span></Link></div>
            <div><span className=' cursor-pointer hover:opacity-80'>Quản lý giao vận</span></div>
            <div><span className=' cursor-pointer'><Avatar /></span></div>
        </div>
    )
}

export default Header