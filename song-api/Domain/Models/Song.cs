namespace Songs.API.Domain.Models
{
    public class Song
    {
        public int Id { get; set; }
        public string Title { get; set; }
		public string Artist { get; set; }
        public int Year { get; set; }
    }
}