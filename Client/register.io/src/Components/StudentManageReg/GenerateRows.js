export function generateRows({
    columnValues = [
        { name: 'coursecode', title: 'Course Code' },
        { name: 'coursenumber', title: 'Course Number' },
        { name: 'coursename', title: 'Course Name' },
        { name: 'credits', title: 'Credits'}
        ],
    length,
    random = 1,
  }) {
    const data = [];
    const columns = Object.keys(columnValues);
  
    for (let i = 0; i < length; i += 1) {
      const record = {};
  
      columns.forEach((column) => {
        let values = columnValues[column];
  
        if (typeof values === 'function') {
          record[column] = values({ random, index: i, record });
          return;
        }
  
        while (values.length === 2 && typeof values[1] === 'object') {
          values = values[1][record[values[0]]];
        }
  
        const value = values[Math.floor(random() * values.length)];
        if (typeof value === 'object') {
          record[column] = { ...value };
        } else {
          record[column] = value;
        }
      });
  
      data.push(record);
    }
  
    return data;
  }