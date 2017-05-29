#ifndef DB_H_
#define DB_H_

#include <string>

class db {
 public:
  db() {}
  std::string Get(std::string key);

  bool Put(std::string key,
           std::string val);
};

#endif
