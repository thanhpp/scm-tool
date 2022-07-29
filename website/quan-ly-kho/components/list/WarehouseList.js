import React, { useEffect, useState } from 'react'
import Delete from '@mui/icons-material/Delete';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';
import WarehouseIcon from '@mui/icons-material/Warehouse';
import ItemWarehouse from '../Item/ItemWarehouse';
import Pagination from '../Pagination/Pagination';
import { useDispatch, useSelector } from 'react-redux';
import { warehouseSliceActions } from '../../store/warehouseSlice';


function WarehouseList() {
    const warehouseData = useSelector((state) => state.warehouse.warehouse)
    // const [warehouseList, setWarehouseList] = useState(warehouseData)

    const dispatch = useDispatch()


    useEffect(() => {
        const fetchWarehouse = async () => {
            console.log(process.env.API_URL)
            try {
                const response = await fetch(`https://scm-tool.thanhpp.ninja/storage`)
                if (!response) {
                    throw new Error('somethign wrong');
                    return;
                }
                const data = await response.json()
                dispatch(warehouseSliceActions.saveWarehouse({
                    data: data.data
                }))
                // setWarehouseList(data)
                // console.log(warehouseData)
            } catch (err) {
                console.log('fetch wrong')
            }
        }

        fetchWarehouse()
    }, [])

    const removeHandle = () => {

    }
    return (
        //bg-[#323259]
        <div className='bg-white rounded-md'>
            <div className='grid grid-cols-12 gap-3 text-black'>
                <div className='flex items-center col-span-1 my-[16px]'>
                    {/* <input type='checkbox' className='mr-[10px] leading-[24px] ml-[20px]' /> */}
                    <WarehouseIcon className='mr-[10px] leading-[24px] ml-[20px]' />
                </div>
                <div className=' col-span-3 my-auto   truncate'>name</div>
                <div className=' col-span-5 my-auto truncate'>description</div>
                <div className=' col-span-2 my-auto truncate'>location</div>
                <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                    <span className=' mr-[10px] rounded-sm'>Delete</span>
                    <span className=' mr-[10px] rounded-sm'>Edit</span>
                </div>
            </div>
            {/* {pagination.currentData ? pagination.currentData.map((item) => <ProductItem isSelectedAll={isSelectedAll} key={item.id} id={item.id} sku={item.sku} name={item.name} category={item.category} price={item.price} amount={item.amount} vendor={item.vendor} arrivalDate={item.arrivalDate} />) : <Box sx={{ display: 'flex' }}> */}
            {warehouseData.map(warehouse => <ItemWarehouse id={warehouse.id} key={warehouse.id} name={warehouse.name} desc={warehouse.desc} location={warehouse.location} />)}
            {/* <CircularProgress /> */}
            {/* </Box>} */}
            {/* <Pagination getNumberPerPage={numberPerPageHandle} pageCount={pagination.pageCount} onPageChange={handleClick} /> */}
            {/* <Pagination /> */}
        </div>
    )
}

export default WarehouseList