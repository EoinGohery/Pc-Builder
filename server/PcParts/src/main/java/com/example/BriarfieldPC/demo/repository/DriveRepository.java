package com.example.BriarfieldPC.demo.repository;

import com.example.BriarfieldPC.demo.entity.Drive;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface DriveRepository extends JpaRepository<Drive,Long> {
    List<Drive> findAll();
}
