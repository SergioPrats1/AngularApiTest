using Songs.API.Domain.Models;

namespace Songs.API.Domain.Services.Communication
{
    public class SongResponse : BaseResponse<Song>
    {
        /// <summary>
        /// Creates a success response.
        /// </summary>
        /// <param name="song">Saved category.</param>
        /// <returns>Response.</returns>
        public SongResponse(Song song) : base(song)
        { }

        /// <summary>
        /// Creates am error response.
        /// </summary>
        /// <param name="message">Error message.</param>
        /// <returns>Response.</returns>
        public SongResponse(string message) : base(message)
        { }
    }
}