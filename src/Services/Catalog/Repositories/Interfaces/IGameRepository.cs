using Catalog.Entities;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Catalog.Repositories.Interfaces {
public interface IGameRepository {
  Task<IEnumerable<Game>> GetGames();
  Task<Game> GetGame(string id);
  Task<IEnumerable<Game>> GetGameByName(string name);
  Task<IEnumerable<Game>> GetGameByCategory(string categoryName);

  Task CreateGame(Game game);
  Task<bool> UpdateGame(Game game);
  Task<bool> DeleteGame(string id);
}
}
