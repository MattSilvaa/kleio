const BASE_URL = 'http://localhost:8080'

export const fetchJobCount = async ({date, location, jobType}:{date: string, location: string, jobType: string}): Promise<number> => {
    const queryParams = new URLSearchParams({date, location, jobType}).toString()
    const url = `${BASE_URL}/api/job-count?${queryParams}`
    try {
        const response = await fetch(url)
        if (!response.ok){
            throw new Error('Network response was not ok')
        }

        const jobCount = await response.json()
        console.log(jobCount)
        return jobCount
    } catch (error){
        console.log('Error fetching jobCount:\t', error)
        return -1
    }
}