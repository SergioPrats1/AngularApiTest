using System.ComponentModel.DataAnnotations;

namespace Songs.API.Resources
{
    public class SaveSongResource
    {
        [Required]
        [MaxLength(80)]
        public string Title { get; set; }
		
        [Required]
        [MaxLength(70)]
        public string Artist { get; set; }		
    }
}