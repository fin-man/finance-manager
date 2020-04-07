export function FormatDate(startDate){
    return startDate.getFullYear() + '-' + ('0' + (startDate.getMonth()+1)).slice(-2) + '-' + ('0' + startDate.getDate()).slice(-2)
};


export default FormatDate;


