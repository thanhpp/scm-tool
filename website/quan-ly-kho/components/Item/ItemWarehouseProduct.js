import React from 'react'
import PowerSettingsNew from '@mui/icons-material/PowerSettingsNew';
import Delete from '@mui/icons-material/Delete';
import WarehouseIcon from '@mui/icons-material/Warehouse';
import Link from 'next/link';
import EditIcon from '@mui/icons-material/Edit';

function ItemWarehouseProduct() {
    return (
        <div className=' grid grid-cols-12 gap-3 text-black  border-black border-t'>
            <div className='flex items-center col-span-1 my-[16px]'>
                {/* <input type='checkbox' className='mr-[10px] leading-[24px] ml-[20px]' /> */}
                <WarehouseIcon className='mr-[10px] leading-[24px] ml-[20px]' />
            </div>
            <Link href='/warehouse-management/hanoi-soc/product/id'><div className=' col-span-3 my-auto hover:border-b border-black cursor-pointer  truncate'>Cây lau nhà</div></Link>
            <div className=' col-span-5 my-auto truncate'>Chổi lau nhà mini brand Lock and Lock là một trong các lựa chọn tốt nhất cho gia đình. Không những giá tiền rẻ sản phẩm cũng có chất lượng khá tốt. Chỉ với khoảng 4 đến 5 trăm là bạn có khả năng có được bộ lau nhà chính hãng mini nhỏ gọn cho gia đình.</div>
            <div className=' col-span-2 my-auto truncate'> Hà Nội</div>
            <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                <span className=' mr-[10px] rounded-sm cursor-pointer hover:opacity-80'><Delete className='' /></span>
                <span className=' mr-[10px] rounded-sm cursor-pointer hover:opacity-80'><EditIcon className='' /></span>
            </div>
        </div>
    )
}

export default ItemWarehouseProduct