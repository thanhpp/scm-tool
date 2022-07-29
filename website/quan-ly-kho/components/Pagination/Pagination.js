import React from 'react'
import ReactPaginate from 'react-paginate'

function Pagination() {

    // const {onPageChange,pageCount,getNumberPerPage} = props
    const handlePageClick = (e) => {
        //   onPageChange(e)
    }

    const getValueHandle = (e) => {
        //   getNumberPerPage(e.target.value)
    }

    return (
        <div className='flex bg-white   p-[20px] border-t border-black'>
            <ReactPaginate
                previousLabel={'<<'}
                nextLabel={'>>'}
                breakLabel={'...'}
                pageCount={''}
                marginPagesDisplayed={2}
                pageRangeDisplayed={5}
                onPageChange={handlePageClick}
                containerClassName={'inline-flex -space-x-px'}
                pageClassName={''}
                pageLinkClassName={'py-2 px-3 ml-0 leading-tight text-gray-500 bg-white  border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white'}
                previousLinkClassName={' py-2 px-3 leading-tight text-gray-500 bg-white rounded-l-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white'}
                nextLinkClassName={' py-2 px-3 leading-tight text-gray-500 bg-white rounded-r-lg border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white'}
            />
            <div className='ml-[5px] '>
                <select onClick={getValueHandle} defaultValue='25' className='bg-[#1e2835] text-white rounded-md'>
                    <option value="25" >25</option>
                    <option value="50">50</option>
                    <option value="75">75</option>
                    <option value="100">100</option>
                </select>
            </div>
        </div>
    )
}

export default Pagination